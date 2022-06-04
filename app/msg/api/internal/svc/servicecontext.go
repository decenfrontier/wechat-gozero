package svc

import (
	"github.com/wslynn/wechat-gozero/app/msg/api/internal/config"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/messageclient"
	"github.com/wslynn/wechat-gozero/app/group/rpc/groupclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	MsgRpc messageclient.MessageClient
	GroupRpc groupclient.GroupClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MsgRpc: messageclient.NewMessageClient(zrpc.MustNewClient(c.MsgRpc)),
		GroupRpc: groupclient.NewGroupClient(zrpc.MustNewClient(c.GroupRpc)),
	}
}
