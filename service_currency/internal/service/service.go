package service

import (
	"go.uber.org/zap"
	v1 "microservices_demo/service_currency/api/v1"
	"microservices_demo/service_currency/internal/biz"

)

type CurrencyService struct {
	v1.UnimplementedCurrencyServiceServer
	shipping *biz.CurrencyUseCase
	logger   *zap.Logger
}

func NewShippingService(shipping *biz.CurrencyUseCase, logger *zap.Logger) *CurrencyService {
	return &CurrencyService{
		shipping: shipping,
		logger:   logger.Named("service_currency"),
	}
}
