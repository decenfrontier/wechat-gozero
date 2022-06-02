package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/wslynn/wechat-gozero/app/msg/model"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/common/xerr"
	"github.com/wslynn/wechat-gozero/proto/msg"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLogic {
	return &PullLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PullLogic) Pull(in *msg.PullRequest) (*msg.PullResponse, error) {
	userId := in.UserId
	groupId := in.GroupId
	platform := in.Platform
	maxMsgId := in.MaxMsgId // 最新的msgId
	// 上次收到在线消息的位置, 没有则为0
	var lastMsgId int64
	// 先查询缓存的last_msg_id
	cache_key := fmt.Sprintf("%d:%s:%s", userId, platform, groupId)
	cache_val, err := l.svcCtx.RedisClient.Get(cache_key)
	if err == nil { // 若查到上次收到在线消息的位置
		lastMsgId, err = strconv.ParseInt(cache_val, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_ERROR),
				"pull cache error, val:%s, err:%v", cache_val, err)
		}
		fmt.Println("last_msg_id:", lastMsgId)
		if lastMsgId >= maxMsgId {
			// 上次收到在线消息的位置 已经是 最新位置, 则不需要拉取
			return nil, nil
		}
	}
	// 查询数据库, 从(maxMsgId, last_msg_id)倒序返回10条数据
	list, err := l.svcCtx.ChatMsgModel.FindMsgListByLastMsgId(l.ctx, groupId, lastMsgId, maxMsgId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
			"message pull error, uid:%s,err:%v", userId, err)
	}
	// 组装返回值, 这里用make(xx, 0)没用, omitempty会自动忽略nil和长度为0的切片
	var resp []*msg.ChatMsg
	if len(list) > 0 {
		for _, chatMsg := range list {
			var pbChatMsg msg.ChatMsg
			_ = copier.Copy(&pbChatMsg, chatMsg)
			pbChatMsg.CreateTime = chatMsg.CreateTime.UnixMilli()
			resp = append(resp, &pbChatMsg)
		}
	}
	return &msg.PullResponse{
		List: resp,
	}, nil
}