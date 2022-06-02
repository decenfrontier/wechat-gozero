package svc

import (
	"github.com/wslynn/wechat-gozero/app/user/api/internal/config"
	"github.com/wslynn/wechat-gozero/app/user/rpc/userclient"

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
