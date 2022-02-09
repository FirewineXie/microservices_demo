package service

import (
	"go.uber.org/zap"
	v1 "microservices_demo/service_email/api/v1"
	"microservices_demo/service_email/internal/biz"
)

type EmailService struct {
	v1.UnimplementedEmailServiceServer
	email  *biz.EmailUseCase
	logger *zap.Logger
}

func NewEmailService(email *biz.EmailUseCase, logger *zap.Logger) *EmailService {
	return &EmailService{
		email:  email,
		logger: logger.Named("service_email"),
	}
}
