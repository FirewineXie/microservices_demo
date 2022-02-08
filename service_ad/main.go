package main

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"microservices_demo/service_ad/internal/server"
	"microservices_demo/service_ad/internal/service"
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
	jaeger, err := jaegerc.InitGlobalTracerProd(&jaegerc.TraceConf{
		ServerName: "service-ad",
	}, logger)
	if err != nil {
		panic(err)
		return
	}
	defer jaeger.Close()
	productService := service.NewAdService(logger)
	go func() {
		addr := "0.0.0.0:" + fmt.Sprint(9001)
		grpcServer = server.NewGRPCServer(logger, productService)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("port is used : %v", err)
		}
		fmt.Println("started grpc server " + addr)
		reflection.Register(grpcServer)
		grpc_prometheus.Register(grpcServer)
		// Register Prometheus metrics handler.
		http.Handle("/metrics", promhttp.Handler())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("start grpc failed :%v", err)
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("deregister service")
	grpcServer.GracefulStop()

}
