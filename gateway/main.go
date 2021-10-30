package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"microservices_demo/gateway/internal/server"
	"microservices_demo/gateway/internal/service"
	"microservices_demo/third_party/jaegerc"
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

func JaegerCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//tracer := opentracing.GlobalTracer()
		spanFromContext, ctx := opentracing.StartSpanFromContext(c.Request.Context(), c.Request.RequestURI)

		c.Set("span", spanFromContext)
		c.Set("ctx", ctx)
		defer spanFromContext.Finish()
		c.Next()

	}
}

func main() {

	var httpServer *http.Server
	engine := gin.Default()

	engine.Use(JaegerCors())
	engine.Use(GinCors())

	jaeger, err := jaegerc.InitGlobalTracerProd(&jaegerc.TraceConf{
		ServerName: "gateway",
	}, logger)
	if err != nil {
		panic(err)
		return
	}
	defer jaeger.Close()
	// 1. 初始化repo
	gatewayService := service.NewGatewayService(logger)

	gatewayService.Router(engine.Group("/api"))

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
