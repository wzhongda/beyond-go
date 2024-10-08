package main

import (
	"beyond-go/pkg/xcode"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"beyond-go/application/article/api/internal/config"
	"beyond-go/application/article/api/internal/handler"
	"beyond-go/application/article/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(xcode.ErrHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
