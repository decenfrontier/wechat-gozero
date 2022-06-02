package svc

import (
	modelGroup "github.com/wslynn/wechat-gozero/app/group/model"
	modelMsg "github.com/wslynn/wechat-gozero/app/msg/model"
	modelUser "github.com/wslynn/wechat-gozero/app/user/model"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      modelUser.UserModel
	GroupModel     modelGroup.GroupModel
	GroupUserModel modelGroup.GroupUserModel
	ChatMsgModel   modelMsg.ChatMsgModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Db.DataSource)
	return &ServiceContext{
		Config:         c,
		UserModel:      modelUser.NewUserModel(conn, c.CacheRedis),
		GroupModel:     modelGroup.NewGroupModel(conn, c.CacheRedis),
		GroupUserModel: modelGroup.NewGroupUserModel(conn, c.CacheRedis),
		ChatMsgModel:   modelMsg.NewChatMsgModel(conn, c.CacheRedis),
	}
}
