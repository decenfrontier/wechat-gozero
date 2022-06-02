package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"ws_chat/app/message/api/internal/svc"
	"ws_chat/app/message/api/internal/types"
	"ws_chat/common/biz"
	"ws_chat/common/ctxdata"
	"ws_chat/common/xerr"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

var lock sync.Mutex

// 存放groupId和group的映射关系
var groupMap = make(map[string]*Group)
// 存放group下的所有在线用户, groupId:[]*Client
var groupUserMap sync.Map

const (
	// 向客户端写入数据的超时时间
	writeWait = 10 * time.Second

	// 接收来自客户端的pong心跳响应包的超时时间
	pongWait = 60 * time.Second

	// 每隔多少秒发送一次ping心跳包
	pingPeriod = (pongWait * 9) / 10

	// 每条消息的最大字节数
	maxMessageSize = 512

	// 最多缓存 待发送的256条消息
	bufSize = 256
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 群组
type Group struct {
	// groupId
	id string
	// 在线的客户端
	onlineClients map[*Client]bool
	// 要广播的群消息
	broadcast chan []byte
	// 保存 刚刚上线 的客户端(处理 进入群聊 事件)
	onEnter chan *Client
	// 保存 刚刚离线 的客户端(处理 离开群聊 事件)
	onLeave chan *Client
}

func GetInstanceGroup(groupId string) *Group {
	group := groupMap[groupId]
	// 双重检 + 互斥锁, 确保多个goroutine同时访问时不会创建多个该Group的实例对象
	if group == nil {
		lock.Lock()
		defer lock.Unlock()
		if groupMap[groupId] == nil {
			// 开始创建实例时, 一般是该群有新消息上传
			group = &Group{
				id: groupId,
				onlineClients: make(map[*Client]bool),
				broadcast:     make(chan []byte),
				onEnter:       make(chan *Client),
				onLeave:       make(chan *Client),
			}
			// 把所有在该群的在线用户都加入该群
			onlineClients, ok := groupUserMap.Load(groupId)
			if ok {
				for _, client := range onlineClients.([]*Client) {
					group.onEnter <- client
				}
			}
			go group.Run()  // 只在创建时运行, 可以保证只运行一次
		}
	}
	return group
}

func (g *Group) Run() {
	for {
		select {
		case client := <-g.onEnter:
			fmt.Printf("group handle onEnter, client:%+v\n", client)
			g.onlineClients[client] = true
			onlineClients, ok := groupUserMap.Load(g.id)
			if !ok {
				onlineClients = []*Client{}
			}
			onlineClients = append(onlineClients.([]*Client), client)
			groupUserMap.Store(g.id, onlineClients)
		case client := <-g.onLeave:
			if _, ok := g.onlineClients[client]; ok {
				delete(g.onlineClients, client)
				close(client.onSend)
				// 群聊中没有人在线, 
				if len(g.onlineClients) == 0 {
					return
				}
			}
		case message := <-g.broadcast:
			for client := range g.onlineClients {
				select {
				case client.onSend <- message:
					fmt.Println("推送消息给客户端1")
				default:
					fmt.Println("客户端缓存满了, 丢弃消息")
					close(client.onSend)
					delete(g.onlineClients, client)
				}
			}
		}
	}
}

// 客户端在服务端的代表
type Client struct {
	// id+platform
	idPlatform string

	// 客户端的所有群, key是groupId
	groupMap map[string]*Group

	// websocket 连接对象
	conn *websocket.Conn

	// 消息数组, 待发送给ws连接的真正客户端
	onSend chan []byte
}

// 从MQ中取出客户端上传的消息, 放到该群的广播队列中
func (c *Client) readPump(svc *svc.ServiceContext) {
	fmt.Println("readPump")
	defer func() {
		for _, group := range c.groupMap {
			group.onLeave <- c
		}
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	// 从MQ中取出消息
	q, err := kq.NewQueue(svc.Config.MqConf, kq.WithHandle(func(key, message string) error {
		logx.Info("从MQ中取出消息:", key, message)
		// 先反序列化, 取出里面的groupId
		msgBytes := []byte(message)
		var chatMsg types.ChatMsg
		err := json.Unmarshal(msgBytes, &chatMsg)
		if err != nil {
			logx.Errorf("【RPC-SRV-ERR】json.Unmarshal failed, message:%s, err: %+v", message, err)
		}
		groupId := chatMsg.GroupId
		// 再根据groupId找到group, 对group进行广播
		group := GetInstanceGroup(groupId)
		group.broadcast <- msgBytes
		return nil
	}))
	if err != nil {
		logx.Errorf("【RPC-SRV-ERR】kq.NewQueue failed c.groupMap:%+v, err: %+v", c.groupMap, err)
	}
	defer func() {
		fmt.Printf("取消订阅, c.groupMap:%+v\n", c.groupMap)
		q.Stop()
	}()
	q.Start()
	fmt.Printf("开始订阅, c.groupMap:%+v\n", c.groupMap)
}

// 向客户端的wsConn中写入数据
func (c *Client) writePump() {
	fmt.Println("writePump")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.onSend:  // 发送消息
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			fmt.Println("推送消息给客户端2")
			w.Write(message)

			// 顺便把所有的消息都发送出去
			for i := 0; i < len(c.onSend); i++ {
				w.Write(<-c.onSend)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:  // 每隔一段时间向客户端发送一个心跳包
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 处理客户端连接
func ServeWs(svc *svc.ServiceContext, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("ServeWs")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return xerr.NewErrCodeMsg(xerr.WS_ERROR, "websocket upgrade failed")
	}
	// 用户默认加入系统通知群组(以0_uid标识), 其它群组等有消息发来再创建
	ctx := r.Context()
	platform, ok := ctx.Value("platform").(string)
	if !ok {
		return xerr.NewErrCodeMsg(xerr.WS_ERROR, "platform is not string")
	}
	uid := ctxdata.GetUidFromCtx(ctx)
	groupId := biz.GetGroupId(0, uid)
	group := GetInstanceGroup(groupId)
	idPlatform := fmt.Sprintf("%d_%s", uid, platform)
	client := &Client{
		idPlatform: idPlatform,
		groupMap: map[string]*Group{groupId: group},
		conn:     conn,
		onSend:   make(chan []byte, bufSize),
	}

	logx.Infof("客户端连接, client:%+v",client)
	// 用户进入群组
	group.onEnter <- client

	// 开启读取和写入协程
	go client.writePump()
	go client.readPump(svc)
	return nil
}

func MqConsumeHandler(key, value string) error {
	log.Println("key:", key, "value:", value)
	return nil
}
