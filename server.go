package main

import (
	"flag"
	"fmt"
	"net"

	mgrpc "intrajasa-merchant-api-gateway/core/grpc"
	log "github.com/Sirupsen/logrus"
	pb "intrajasa-merchant-api-gateway/core/grpc/services"
	
	"intrajasa-merchant-api-gateway/core/router"
	"google.golang.org/grpc"
)

var (
	env = flag.String("env", "development", "running environment")
)

// Api server start from here. router is define your api router and public it.
func main() {
	flag.Parse()

	// GRPC
	//
	// Here will enable grpc server, if you don`t want it, you can disable it
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 10000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		pb.RegisterRouteGuideServer(grpcServer, mgrpc.NewServer())
		grpcServer.Serve(lis)
	}()

	// HTPP
	//
	// start api server, *env is what`s environment will running, currentlly this only for enable or disable debug modle
	// After may be use it load different varible.
	router.Start(*env)
}
