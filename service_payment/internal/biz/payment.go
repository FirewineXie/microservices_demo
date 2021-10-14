package biz

import (
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type PaymentUseCase struct {
	logger *zap.Logger
}

// CardValidator 银行卡验证，这里简单化，
// 前缀时 6546 开头
// 长度 15
func (c PaymentUseCase) CardValidator(number string) (cardType string, valid bool) {
	_, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return "", false
	}

	if len(number) > 18 && len(number) < 0 {
		return "", false
	}
	if strings.HasPrefix(number, "6546") {
		return "visa", true
	}
	if strings.HasPrefix(number, "6546") {
		return "mastercard", true
	}

	return "other", true
}

func NewPaymentUseCase(logger *zap.Logger) *PaymentUseCase {
	return &PaymentUseCase{
		logger: logger,
	}
}
