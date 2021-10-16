package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"microservices_demo_v1/gateway/internal/server"
	"microservices_demo_v1/gateway/internal/service"
	"time"

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

	var httpServer *http.Server
	engine := gin.Default()

	// 1. 初始化repo

	gatewayService := service.NewGatewayService(logger)

	gatewayService.Router(engine.Group(""))
	go func() {

		addr := "0.0.0.0:8000"
		httpServer = server.NewHTTPServer(logger, engine, addr, 5000)
		fmt.Println("started http server" + addr)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("start http failed : %v", err)
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("deregister service")
	stopService(httpServer)
}

func stopService(httpServer *http.Server) {

	log.Println("shutdown Server ....")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case <-ctx.Done():
		// wait data
		log.Println("Server exiting")
	}
	log.Println("Server exiting")
}
