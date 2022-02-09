package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"microservices_demo/service_shipping/internal/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9009", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	//resp, err := v1.NewShippingServiceClient(conn).ShipOrder(ctx, &v1.ShipOrderRequest{})
	//if err != nil {
	//	fmt.Printf("could not charge the card: %+v", err)
	//	return
	//}
	//fmt.Println(resp.GetTrackingId())

	resp1, err1 := v1.NewShippingServiceClient(conn).GetQuote(ctx, &v1.GetQuoteRequest{})
	if err1 != nil {
		fmt.Printf("could not charge the card: %+v", err)
		return
	}
	fmt.Println(resp1.GetCostUsd())
}
