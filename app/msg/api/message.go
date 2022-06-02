package main

import (
	"flag"
	"fmt"
	"net/http"

	"ws_chat/app/message/api/internal/config"
	"ws_chat/app/message/api/internal/handler"
	"ws_chat/app/message/api/internal/logic"
	"ws_chat/app/message/api/internal/svc"
	"ws_chat/common/xresp"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	// 禁用显示cpu
	logx.DisableStat()
	// http升级为websocket
	server.AddRoute(
		rest.Route{
			Method: http.MethodGet,
			Path:   "/ws",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				err := logic.ServeWs(ctx, w, r)
				xresp.Response(r, w, nil, err)
			},
		},
		rest.WithJwt(ctx.Config.JwtAuth.AccessSecret),
	)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
