package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_email/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9001", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	paymentResp, err := v1.NewEmailServiceClient(conn).SendOrderConfirmation(ctx, &v1.SendOrderConfirmationRequest{
		Email: "12334",
	})
	if err != nil {
		fmt.Printf("could not charge the card: %+v", err)
		return
	}
	fmt.Println(paymentResp.String())
}
