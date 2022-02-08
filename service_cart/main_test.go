package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"microservices_demo/service_cart/internal/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx,
		"0.0.0.0:9002", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	cart, _ := v1.NewCartServiceClient(conn).AddItem(ctx, &v1.AddItemRequest{
		UserId: "1",
		Item: &v1.CartItem{
			ProductId: "t1",
			Quantity:  123,
		},
	})
	fmt.Println(cart.String())
}
