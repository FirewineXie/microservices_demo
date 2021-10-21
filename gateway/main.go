package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"microservices_demo/gateway/internal/server"
	"microservices_demo/gateway/internal/service"
	"net/http"
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
	engine.Use(GinLogger())
	// 1. 初始化repo
	//gatewayService := service.NewGatewayService(logger)
	group := engine.Group("")
	{
		group.GET("/product/:id", service.ProductHandler)
		group.GET("/cart", service.ViewCartHandler)
		group.POST("/cart", service.AddToCartHandler)
		group.POST("/cart/empty", service.EmptyCartHandler)
		group.POST("/setCurrency", service.SetCurrencyHandler)
		group.GET("/logout", service.LogoutHandler)
		group.POST("/cart/checkout", service.PlaceOrderHandler)
	}

	addr := ":8000"
	httpServer = server.NewHTTPServer(logger, engine, addr, 5000)
	fmt.Println("started http server" + addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("start http failed : %v", err)
	}
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//fmt.Println("deregister service")
	//stopService(httpServer)
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

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		fmt.Println(c.Writer.Status())
		fmt.Println(c.Writer.Write([]byte("44234")))
		fmt.Println(c.Writer.WriteString("4543534"))
		fmt.Printf("url=%s, status=%d, resp=%s \n", c.Request.URL, c.Writer.Status(), blw.body.String())
	}
}
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}