package main

import (
	"flag"
	"fmt"
	"net/http"
	"swap-api/internal/config"
	"swap-api/internal/handler"
	"swap-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	// 日志的打印格式替换一下
	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token,x-auth-token")
	}, nil, "*"))
	defer server.Stop()

	// 注册路由
	router := handler.NewRouters(server)
	handler.RegisterHandlers(router, ctx)

	fmt.Printf("Starting http server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
