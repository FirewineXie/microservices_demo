package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo_v1/service_payment/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9002", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	paymentResp, err := v1.NewPaymentServiceClient(conn).Charge(ctx, &v1.ChargeRequest{
		Amount: &v1.Money{
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
