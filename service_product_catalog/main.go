package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"microservices_demo/service_product_catalog/internal/biz"
	"microservices_demo/service_product_catalog/internal/pkg"
	"microservices_demo/service_product_catalog/internal/server"
	"microservices_demo/service_product_catalog/internal/service"

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

	useCase := biz.NewProductCatalogUseCase(logger)
	productService := service.NewProductCatalogService(useCase, logger)
	pkg.ConnectNacos()
	_, err := pkg.RegisterInstance()
	if err != nil {
		panic(err)
	}
	go func() {
		addr := "0.0.0.0:" + fmt.Sprint(9007)
		grpcServer = server.NewGRPCServer(logger, productService)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Error("port is used : ", zap.Error(err))
			return
		}
		fmt.Println("started grpc server" + addr)
		reflection.Register(grpcServer)
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("start grpc failed :%v", zap.Error(err))
			return
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	pkg.DeregisterInstance()
	fmt.Println("deregister service")
	grpcServer.GracefulStop()
}
