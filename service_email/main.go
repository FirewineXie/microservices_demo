package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"microservices_demo/service_email/internal/biz"
	"microservices_demo/service_email/internal/server"
	"microservices_demo/service_email/internal/service"
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
	//tracer := startTracer()
	//defer tracer.Close()
	emailUseCase := biz.NewEmailUseCase(logger)
	productService := service.NewEmailService(emailUseCase, logger)
	//pkg.ConnectNacos()
	//_, err := pkg.RegisterInstance()
	//if err != nil {
	//	panic(err)
	//}
	go func() {
		addr := "0.0.0.0:" + fmt.Sprint(9005)
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
	//pkg.DeregisterInstance()
	fmt.Println("deregister service")
	grpcServer.GracefulStop()
}
func startTracer() io.Closer {
	closer, err := jaegerc.InitGlobalTracerProd(&jaegerc.TraceConf{

		ServerName: "service_email",
	}, logger)
	if err != nil {
		panic("Could not initialize jaeger tracer: %s" + err.Error())

	}
	return closer
}
