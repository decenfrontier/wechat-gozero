package svc

import (
	"github.com/wslynn/wechat-gozero/app/group/api/internal/config"
	"github.com/wslynn/wechat-gozero/app/group/rpc/groupclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	GroupRpc groupclient.GroupClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GroupRpc: groupclient.NewGroupClient(zrpc.MustNewClient(c.GroupRpc)),
	}
}
