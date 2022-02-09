package biz

import (
	"context"
	"go.uber.org/zap"
)

type EmailUseCase struct {
	logger *zap.Logger
}

func (c EmailUseCase) SendOrderResultByEmail(ctx context.Context, email string, order string) error {
	c.logger.Info("发送邮件成功：", zap.String("email", email), zap.String("data", order))
	return nil
}

func NewEmailUseCase (logger *zap.Logger) *EmailUseCase{
	return &EmailUseCase{
		logger: logger,
	}
}
