package svc

import (
	"time"

	"github.com/segmentio/kafka-go"
	modelGroup "github.com/wslynn/wechat-gozero/app/group/model"
	"github.com/wslynn/wechat-gozero/app/group/rpc/internal/config"
	modelMsg "github.com/wslynn/wechat-gozero/app/msg/model"
	modelUser "github.com/wslynn/wechat-gozero/app/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	GroupModel     modelGroup.GroupModel
	GroupUserModel modelGroup.GroupUserModel
	UserModel      modelUser.UserModel
	ChatMsgModel   modelMsg.ChatMsgModel
	MqWriter     *kafka.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)
	mqWriter := &kafka.Writer{
		Addr:     kafka.TCP(c.MqConf.Brokers...),
		Topic:    c.MqConf.Topic,
		BatchTimeout: time.Millisecond * 20,
	}
	return &ServiceContext{
		Config:         c,
		GroupModel:     modelGroup.NewGroupModel(sqlConn, c.Cache),
		GroupUserModel: modelGroup.NewGroupUserModel(sqlConn, c.Cache),
		UserModel:      modelUser.NewUserModel(sqlConn, c.Cache),
		ChatMsgModel:   modelMsg.NewChatMsgModel(sqlConn, c.Cache),
		MqWriter: mqWriter,
	}
}
