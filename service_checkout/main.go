package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"microservices_demo/service_checkout/internal/biz"
	"microservices_demo/service_checkout/internal/server"
	"microservices_demo/service_checkout/internal/service"
	"microservices_demo/third_party/jaegerc"

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
	tracer := startTracer()
	defer tracer.Close()
	useCase := biz.NewCheckoutUseCase(logger)
	productService := service.NewCheckoutService(useCase, logger)

	go func() {
		addr := "0.0.0.0:" + fmt.Sprint(9003)
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

	fmt.Println("deregister service")
	grpcServer.GracefulStop()
}

func startTracer() io.Closer {
	closer, err := jaegerc.InitGlobalTracerProd(&jaegerc.TraceConf{

		ServerName: "service_checkout",
	}, logger)
	if err != nil {
		panic("Could not initialize jaeger tracer: %s" + err.Error())

	}
	return closer
}
