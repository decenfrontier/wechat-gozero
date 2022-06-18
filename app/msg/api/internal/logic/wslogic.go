package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/wslynn/wechat-gozero/app/msg/api/internal/svc"
	"github.com/wslynn/wechat-gozero/app/msg/api/internal/types"
	"github.com/wslynn/wechat-gozero/common/biz"
	"github.com/wslynn/wechat-gozero/common/ctxdata"
	"github.com/wslynn/wechat-gozero/common/xerr"
	pbgroup "github.com/wslynn/wechat-gozero/proto/group"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var lock sync.Mutex

// 存放groupId对应的group, groupId:*Group
var globalGroupMap = make(map[string]*Group)

const (
	// 向客户端写入数据的超时时间
	writeWait = 10 * time.Second

	// 接收来自客户端的pong心跳响应包的超时时间
	pongWait = 60 * time.Second

	// 每隔多少秒发送一次ping心跳包
	pingPeriod = (pongWait * 9) / 10

	// 每条消息的最大字节数
	maxMessageSize = 512

	// 最多缓存 待发送的10条消息
	bufSize = 10
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
	group := globalGroupMap[groupId]
	// 双重检 + 互斥锁, 确保多个goroutine同时访问时不会创建多个该Group的实例对象
	if group == nil {
		lock.Lock()
		defer lock.Unlock()
		group = globalGroupMap[groupId]
		if group == nil {
			group = &Group{
				id:            groupId,
				onlineClients: make(map[*Client]bool),
				broadcast:     make(chan []byte),
				onEnter:       make(chan *Client),
				onLeave:       make(chan *Client),
			}
			// 把该群组加到全局map中
			globalGroupMap[groupId] = group
			go group.Run() // 只在创建时运行, 可以保证只运行一次
		}
	}
	return group
}

func (g *Group) Run() {
	for {
		select {
		case client := <-g.onEnter:  // 刚刚上线
			fmt.Printf("group handle onEnter, client:%+v\n", client)
			g.onlineClients[client] = true

		case client := <-g.onLeave:  // 刚刚离线
			delete(g.onlineClients, client)

		case message := <-g.broadcast: // 有新消息上传
			for client := range g.onlineClients {
				select {
				case client.onSend <- message:
					logx.Info("group send message to client")
				default:
					fmt.Println("客户端缓存满了, 可能是连接异常, 让客户端离线")
					g.onLeave <- client
				}
			}
		}
	}
}

// 客户端在服务端的代表
type Client struct {
	idPlatform string  // id+platform
	groupMap map[string]*Group  // 客户端的所有群, key是groupId
	conn *websocket.Conn  // websocket 连接对象
	onSend chan []byte  // 消息数组, 待发送给ws连接的真正客户端
}

// 从MQ中取出消息
func ConsumeMsgFromMQ(svc *svc.ServiceContext) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   svc.Config.MqConf.Brokers,
		Topic:     svc.Config.MqConf.Topic,
		GroupID:   svc.Config.MqConf.Group,
	})
	ctx := context.Background()
	for {
		kfkMsg, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		// 先反序列化, 取出里面的groupId
		msgBytes := kfkMsg.Value
		message := string(msgBytes)
		logx.Info("从mq消费到消息:%s", message)

		var chatMsg types.ChatMsg
		err = json.Unmarshal(msgBytes, &chatMsg)
		if err != nil {
			logx.Errorf("【RPC-SRV-ERR】json.Unmarshal failed, message:%s, err: %+v", message, err)
		}
		groupId := chatMsg.GroupId
		// 再根据groupId找到group, 对group进行广播
		group := GetInstanceGroup(groupId)
		group.broadcast <- msgBytes
		// 处理完后, 再提交offset
		err = r.CommitMessages(ctx, kfkMsg)
		if err != nil {
			logx.Errorf("【RPC-SRV-ERR】CommitMessages failed, err: %+v", err)
		}
	}
	if err := r.Close(); err != nil {
		logx.Error("failed to close reader:", err)
	}
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
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			logx.Infof("【RPC-SRV-INFO】client %s disconnect, reason: %+v", c.idPlatform, err)
			break
		}
		fmt.Println("readPump msg:", string(msg))  // ws消息不做其它处理, 因为消息是通过http上传的
	}
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
		case message, ok := <-c.onSend: // 发送消息
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 顺便把所有的消息都发送出去
			for i := 0; i < len(c.onSend); i++ {
				w.Write(<-c.onSend)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C: // 每隔一段时间向客户端发送一个心跳包
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
	// 获取上下文对象
	ctx := r.Context()
	// 获取 平台号 和 uid
	platform := r.Header.Get("platform")
	if platform == "" {
		return xerr.NewErrCodeMsg(xerr.WS_ERROR, "platform is not string")
	}
	uid := ctxdata.GetUidFromCtx(ctx)
	// 获取用户的所有群组
	resp, err := svc.GroupRpc.UserGroupList(ctx, &pbgroup.UserGroupListRequest{UserId: uid})
	if err != nil {
		return err
	}
	// 构造client
	groupMap := map[string]*Group{}
	idPlatform := fmt.Sprintf("%d_%s", uid, platform)
	client := &Client{
		idPlatform: idPlatform,
		groupMap:   groupMap,
		conn:       conn,
		onSend:     make(chan []byte, bufSize),
	}
	fmt.Printf("客户端连接, client:%+v\n", client)
	// 用户进入群组, 默认加入系统通知群组(以0_uid标识)
	userGroupIdList := append(resp.List, biz.GetGroupId(0, uid))
	for _, groupId := range userGroupIdList {
		group := GetInstanceGroup(groupId)
		groupMap[groupId] = group
		group.onEnter <- client
	}
	// 设置client的groupMap
	client.groupMap = groupMap
	// 开启读取和写入协程
	go client.writePump()
	go client.readPump(svc)
	return nil
}
