package service

import (
	"go.uber.org/zap"
	"microservices_demo/service_payment/internal/api/v1"
	"microservices_demo/service_payment/internal/biz"
)

type PaymentService struct {
	v1.UnimplementedPaymentServiceServer
	payment *biz.PaymentUseCase
	logger  *zap.Logger
}

func NewPaymentService(payment *biz.PaymentUseCase, logger *zap.Logger) *PaymentService {
	return &PaymentService{
		payment: payment,
		logger:  logger.Named("service_payment"),
	}
}
