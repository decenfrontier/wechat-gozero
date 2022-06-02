package svc

import (
	"wechat-gozero/app/message/api/internal/config"
	"wechat-gozero/app/message/rpc/messageclient"

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
