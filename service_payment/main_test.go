package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v12 "microservices_demo/service_payment/internal/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9006", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	paymentResp, err := v12.NewPaymentServiceClient(conn).Charge(ctx, &v12.ChargeRequest{
		Amount: &v12.Money{
			CurrencyCode: "11",
			Units:        222,
			Nanos:        33,
		},
	})
	if err != nil {
		fmt.Printf("could not charge the card: %+v", err)
		return
	}
	fmt.Println(paymentResp.GetTransactionId())
}
