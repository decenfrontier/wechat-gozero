package main

import (
	"flag"
	"fmt"

	"github.com/wslynn/wechat-gozero/app/user/api/internal/config"
	"github.com/wslynn/wechat-gozero/app/user/api/internal/handler"
	"github.com/wslynn/wechat-gozero/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
