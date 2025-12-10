package main

import (
	"flag"
	"fmt"

	"core-service/core"
	"core-service/internal/config"
	emotionserviceServer "core-service/internal/server/emotionservice"
	intentserviceServer "core-service/internal/server/intentservice"
	promptserviceServer "core-service/internal/server/promptservice"
	ragserviceServer "core-service/internal/server/ragservice"
	routingserviceServer "core-service/internal/server/routingservice"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		core.RegisterRagServiceServer(grpcServer, ragserviceServer.NewRagServiceServer(ctx))
		core.RegisterIntentServiceServer(grpcServer, intentserviceServer.NewIntentServiceServer(ctx))
		core.RegisterEmotionServiceServer(grpcServer, emotionserviceServer.NewEmotionServiceServer(ctx))
		core.RegisterRoutingServiceServer(grpcServer, routingserviceServer.NewRoutingServiceServer(ctx))
		core.RegisterPromptServiceServer(grpcServer, promptserviceServer.NewPromptServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
