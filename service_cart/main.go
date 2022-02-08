package main

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"microservices_demo/service_cart/internal/biz"
	"microservices_demo/service_cart/internal/server"
	"microservices_demo/service_cart/internal/service"
	"microservices_demo/third_party/jaegerc"
	"net"
	"net/http"
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
	useCase := biz.NewCartUseCase(logger)
	productService := service.NewCartService(useCase, logger)
	jaeger, err := jaegerc.InitGlobalTracerProd(&jaegerc.TraceConf{
		ServerName: "service-cart",
	}, logger)
	if err != nil {
		panic(err)
		return
	}
	defer jaeger.Close()
	go func() {
		addr := "0.0.0.0:" + fmt.Sprint(9002)
		grpcServer = server.NewGRPCServer(logger, productService)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Error("port is used : ", zap.Error(err))
			return
		}
		fmt.Println("started grpc server" + addr)
		reflection.Register(grpcServer)
		grpc_prometheus.Register(grpcServer)
		http.Handle("/metrics", promhttp.Handler())
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("start grpc failed :%v", zap.Error(err))
			return
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("deregister service")
	grpcServer.GracefulStop()
}
