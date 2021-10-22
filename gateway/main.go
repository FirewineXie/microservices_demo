package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"microservices_demo/gateway/internal/server"
	"microservices_demo/gateway/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger *zap.Logger

func init() {
	logger = zap.NewExample()
}

func GinCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"X-Requested-With", "Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func main() {

	var httpServer *http.Server
	engine := gin.Default()

	engine.Use(GinCors())
	// 1. 初始化repo
	gatewayService := service.NewGatewayService(logger)


	addr := ":8000"
	httpServer = server.NewHTTPServer(logger, engine, addr)
	fmt.Println("started http server" + addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("start http failed : %v", err)
	}

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
