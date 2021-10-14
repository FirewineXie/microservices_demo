package service

import (
	"go.uber.org/zap"
	v1 "microservices_demo_v1/service_shipping/api/v1"

	"microservices_demo_v1/service_shipping/internal/biz"
)

type ShippingService struct {
	v1.UnimplementedShippingServiceServer
	shipping  *biz.ShippingUseCase
	logger *zap.Logger
}

func NewShippingService(shipping *biz.ShippingUseCase, logger *zap.Logger) *ShippingService {
	return &ShippingService{
		shipping:  shipping,
		logger: logger.Named("service_payment"),
	}
}
