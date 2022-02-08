package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"microservices_demo/service_checkout/internal/api/v1"
	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx,
		"0.0.0.0:9003", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	resp, _ := v1.NewCheckoutServiceClient(conn).PlaceOrder(ctx, &v1.PlaceOrderRequest{
		Email:  "3333@173.com",
		UserId: "3232",
	})
	fmt.Println(resp.GetOrder())

}
