package main
import (
	"fmt"
	"net"

	mgrpc "intrajasa-merchant-api-gateway/Core/Grpc"
	log "github.com/Sirupsen/logrus"
	pb "intrajasa-merchant-api-gateway/Core/Grpc/Services"
	
	"intrajasa-merchant-api-gateway/Config"
	"intrajasa-merchant-api-gateway/Database"
	"intrajasa-merchant-api-gateway/Core/Router"
	"intrajasa-merchant-api-gateway/Core/Models"
	"google.golang.org/grpc"
)

// Api server start from here. router is define your api router and public it.
func main() {
	// GORM DATABASE
	database.Mysql, database.Err = database.ConnectToDB("main")
	if database.Err != nil {
		fmt.Println("status error : ", database.Err)
	} else {
		fmt.Println("database connected")
	}
	defer database.Mysql.Close()
	database.Mysql.AutoMigrate(&Models.MerchantVa{})

	// GRPC
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
	app_env := config.GoDotEnvVariable("APP_ENV")

	// HTPP
	// start api server, *env is what`s environment will running, currentlly this only for enable or disable debug modle
	// After may be use it load different varible.
	router.Start(app_env)
}
