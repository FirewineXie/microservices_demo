package server

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewHTTPServer(logger *zap.Logger,addr string, timeout time.Duration) *http.Server {
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: timeout,
	}

	return server
}
