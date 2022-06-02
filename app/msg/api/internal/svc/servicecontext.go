package svc

import (
	"github.com/wslynn/wechat-gozero/app/msg/api/internal/config"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/messageclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	MessageRpc messageclient.MessageClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MessageRpc: messageclient.NewMessageClient(zrpc.MustNewClient(c.MessageRpc)),
	}
}
