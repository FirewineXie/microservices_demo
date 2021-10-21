package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewHTTPServer(logger *zap.Logger, handler *gin.Engine, addr string, timeout time.Duration) *http.Server {
	server := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: timeout,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}
