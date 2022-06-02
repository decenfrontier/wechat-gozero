package logic

import (
	"context"

	"wechat-gozero/app/message/model"
	"wechat-gozero/app/message/rpc/internal/svc"
	"wechat-gozero/app/message/rpc/proto"
	"wechat-gozero/common/xerr"
	"wechat-gozero/common/xmq"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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

func (l *UploadLogic) Upload(in *proto.UploadRequest) (*proto.UploadResponse, error) {
	msg := &model.ChatMsg{
		GroupId:  in.GroupId,
		SenderId: in.SenderId,
		Type:     in.Type,
		Content:  in.Content,
		Uuid:     in.Uuid,
	}

	err := l.svcCtx.ChatMsgModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 存到数据库
		ret, err := l.svcCtx.ChatMsgModel.TransInsert(l.ctx, session, msg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "消息uuid已存在"), "insert message failed, msg: %+v", msg)
		}
		msg.Id, err = ret.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "获取消息id失败"), "get message id failed, msg: %+v", msg)
		}
		dbMsg, err := l.svcCtx.ChatMsgModel.TransFindOne(l.ctx, session, msg.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "获取消息失败"), "get message failed, msg: %+v", msg)
		}
		// 放入消息队列
		var mqMsg proto.ChatMsg
		copier.Copy(&mqMsg, dbMsg)
		mqMsg.CreateTime = dbMsg.CreateTime.UnixMilli()
		msg.CreateTime = dbMsg.CreateTime
		err = xmq.PushToMq(l.svcCtx.MqProducer, &mqMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(xerr.MQ_ERROR, "消息推送失败"), "push message to mq failed, msg: %+v", msg)
		}
		logx.Infof("push to mq msg: %+v", msg)
		// commit
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &proto.UploadResponse{
		Id:         msg.Id,
		CreateTime: msg.CreateTime.UnixMilli(),
	}, nil
}
