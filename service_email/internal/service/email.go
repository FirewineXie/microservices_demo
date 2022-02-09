package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "microservices_demo/service_email/api/v1"
)

func (f *EmailService) SendOrderConfirmation(ctx context.Context,
	request *v1.SendOrderConfirmationRequest) (response *v1.Empty, err error) {
	result := v1.Empty{}
	email := request.GetEmail()
	order := request.GetOrder()

	// 这里相当于生成确认信息
	f.logger.Sugar().Infof("send email to %v of content is %v", email, order.String())

	if err = f.email.SendOrderResultByEmail(ctx, email, order.String()); err != nil {
		return &result, status.Errorf(codes.Internal, "don't send")
	}
	return &result, err
}
