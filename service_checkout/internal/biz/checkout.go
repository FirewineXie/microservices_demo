package biz

import "go.uber.org/zap"

type CheckoutUseCase struct {
	logger *zap.Logger
}

func NewCheckoutUseCase(logger *zap.Logger) *CheckoutUseCase {
	return &CheckoutUseCase{logger}
}
