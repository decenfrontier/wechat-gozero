package svc

import (
	modelGroup "ws_chat/app/group/model"
	modelUser "ws_chat/app/user/model"
	modelMsg "ws_chat/app/message/model"
	"ws_chat/app/group/rpc/internal/config"

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
