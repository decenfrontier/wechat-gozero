package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	MessageRpc zrpc.RpcClientConf
	GroupRpc zrpc.RpcClientConf
	MqConf kq.KqConf
}
