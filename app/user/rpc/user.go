package main

import (
	"flag"
	"fmt"

	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/config"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/server"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/proto/user"
	"github.com/wslynn/wechat-gozero/common/interceptor"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	svr := server.NewUserClientServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserClientServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 添加拦截器
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	// 禁用显示cpu
	logx.DisableStat()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
