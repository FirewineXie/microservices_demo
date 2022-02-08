package service

import (
	"go.uber.org/zap"
	"microservices_demo/service_checkout/internal/api/v1"
	"microservices_demo/service_checkout/internal/biz"
)

type CheckoutService struct {
	v1.UnimplementedCheckoutServiceServer
	checkout *biz.CheckoutUseCase
	logger   *zap.Logger
}

func NewCheckoutService(checkout *biz.CheckoutUseCase, logger *zap.Logger) *CheckoutService {
	return &CheckoutService{
		checkout: checkout,
		logger:   logger,
	}
}
