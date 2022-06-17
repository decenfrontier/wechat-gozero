package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	MsgRpc   zrpc.RpcClientConf
	GroupRpc zrpc.RpcClientConf
	MqConf   struct {
		Brokers []string
		Topic   string
		Group   string
	}
}
