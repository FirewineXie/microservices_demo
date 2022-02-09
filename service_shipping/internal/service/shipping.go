package service

import (
	"go.uber.org/zap"
	"microservices_demo/service_shipping/internal/api/v1"

	"microservices_demo/service_shipping/internal/biz"
)

type ShippingService struct {
	v1.UnimplementedShippingServiceServer
	shipping *biz.ShippingUseCase
	logger   *zap.Logger
}

func NewShippingService(shipping *biz.ShippingUseCase, logger *zap.Logger) *ShippingService {
	return &ShippingService{
		shipping: shipping,
		logger:   logger.Named("service_payment"),
	}
}
