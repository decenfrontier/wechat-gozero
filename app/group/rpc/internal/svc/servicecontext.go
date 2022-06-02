package svc

import (
	modelGroup "wechat-gozero/app/group/model"
	"wechat-gozero/app/group/rpc/internal/config"
	modelMsg "wechat-gozero/app/message/model"
	modelUser "wechat-gozero/app/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	GroupModel modelGroup.GroupModel
	GroupUserModel modelGroup.GroupUserModel
	UserModel modelUser.UserModel
	ChatMsgModel modelMsg.ChatMsgModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)
	return &ServiceContext{
		Config: c,
		GroupModel: modelGroup.NewGroupModel(sqlConn, c.Cache),
		GroupUserModel: modelGroup.NewGroupUserModel(sqlConn, c.Cache),
		UserModel: modelUser.NewUserModel(sqlConn, c.Cache),
		ChatMsgModel: modelMsg.NewChatMsgModel(sqlConn, c.Cache),
	}
}
