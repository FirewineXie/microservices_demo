package service

import (
	"go.uber.org/zap"
)

type GatewayService struct {
	logger *zap.Logger
}



func NewGatewayService(logger *zap.Logger) *GatewayService {
	return &GatewayService{logger}
}
