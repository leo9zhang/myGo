package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	config2 "mall/service/user/api/internal/config"
	handler2 "mall/service/user/api/internal/handler"
	svc2 "mall/service/user/api/internal/svc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config2.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc2.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler2.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
