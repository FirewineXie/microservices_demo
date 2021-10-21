package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_cart/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx,
		"0.0.0.0:9003", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	cart, _ := v1.NewCartServiceClient(conn).GetCart(ctx, &v1.GetCartRequest{
		UserId: "1",
	})
	fmt.Println(cart.GetItems())
}
