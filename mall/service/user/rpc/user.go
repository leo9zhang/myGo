package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	config2 "mall/service/user/rpc/internal/config"
	server2 "mall/service/user/rpc/internal/server"
	svc2 "mall/service/user/rpc/internal/svc"
	user2 "mall/service/user/rpc/user"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config2.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc2.NewServiceContext(c)
	srv := server2.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user2.RegisterUserServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
