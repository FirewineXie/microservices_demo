package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"microservices_demo/service_ad/internal/pkg"
	"microservices_demo/service_ad/internal/server"
	"microservices_demo/service_ad/internal/service"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger *zap.Logger

func init() {
	logger = zap.NewExample()
}
func main() {
	var grpcServer *grpc.Server

	productService := service.NewAdService(logger)
	pkg.ConnectNacos()
	_, err := pkg.RegisterInstance()
	if err != nil {
		panic(err)
	}
	go func() {
		addr := "0.0.0.0:" + fmt.Sprint(9001)
		grpcServer = server.NewGRPCServer(logger, productService)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("port is used : %v", err)
		}
		fmt.Println("started grpc server" + addr)
		reflection.Register(grpcServer)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("start grpc failed :%v", err)
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	pkg.DeregisterInstance()
	fmt.Println("deregister service")
	grpcServer.GracefulStop()

}
