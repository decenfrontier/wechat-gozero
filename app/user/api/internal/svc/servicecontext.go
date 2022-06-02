package svc

import (
	"ws_chat/app/user/api/internal/config"
	"ws_chat/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUserClient(zrpc.MustNewClient(c.UserRpc)),
	}
}
