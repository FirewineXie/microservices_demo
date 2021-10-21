package service

import (
	"go.uber.org/zap"
	v1 "microservices_demo/service_cart/api/v1"
	"microservices_demo/service_cart/internal/biz"
)

type CartService struct {
	v1.UnimplementedCartServiceServer
	cart   *biz.CartUseCase
	logger *zap.Logger
}

func NewCartService(cart *biz.CartUseCase,logger *zap.Logger) *CartService{
	return &CartService{
		cart: cart,
		logger: logger,
	}
}
