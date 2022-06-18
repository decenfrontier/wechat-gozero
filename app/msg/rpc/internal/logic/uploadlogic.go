package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/wslynn/wechat-gozero/app/msg/model"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/common/xerr"
	"github.com/wslynn/wechat-gozero/common/xmq"
	"github.com/wslynn/wechat-gozero/proto/msg"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *msg.UploadRequest) (*msg.UploadResponse, error) {
	chatMsg := &model.ChatMsg{
		GroupId:  in.GroupId,
		SenderId: in.SenderId,
		Type:     in.Type,
		Content:  in.Content,
		Uuid:     in.Uuid,
	}

	err := l.svcCtx.ChatMsgModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 存到数据库
		ret, err := l.svcCtx.ChatMsgModel.TransInsert(l.ctx, session, chatMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "消息uuid已存在"), "insert message failed, msg: %+v", chatMsg)
		}
		chatMsg.Id, err = ret.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "获取消息id失败"), "get message id failed, msg: %+v", chatMsg)
		}
		dbMsg, err := l.svcCtx.ChatMsgModel.TransFindOne(l.ctx, session, chatMsg.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "获取消息失败"), "get message failed, msg: %+v", chatMsg)
		}
		// 放入消息队列
		var mqMsg msg.ChatMsg
		copier.Copy(&mqMsg, dbMsg)
		mqMsg.CreateTime = dbMsg.CreateTime.UnixMilli()
		chatMsg.CreateTime = dbMsg.CreateTime
		err = xmq.PushToMq(l.ctx, l.svcCtx.MqWriter, &mqMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.MQ_ERROR, "消息推送失败"), "push message to mq failed, msg: %+v, err: %v", chatMsg, err)
		}
		logx.Infof("push to mq msg: %+v", chatMsg)
		// commit
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &msg.UploadResponse{
		Id:         chatMsg.Id,
		CreateTime: chatMsg.CreateTime.UnixMilli(),
	}, nil
}
