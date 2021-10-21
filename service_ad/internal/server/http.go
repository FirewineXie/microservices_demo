package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewHTTPServer(logger *zap.Logger, handler *gin.Engine, addr string, timeout time.Duration) *http.Server {
	handler.Use(gin.Logger())
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		WriteTimeout: timeout,
	}

	return server
}
