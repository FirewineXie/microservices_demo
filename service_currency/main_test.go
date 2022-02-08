package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v12 "microservices_demo/service_currency/internal/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9004", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	//resp, err := v12.NewCurrencyServiceClient(conn).Convert(ctx, &v12.CurrencyConversionRequest{})
	//if err != nil {
	//	fmt.Printf("could not charge the card: %+v", err)
	//	return
	//}
	//fmt.Println(resp.GetCurrencyCode())

	resp1, err1 := v12.NewCurrencyServiceClient(conn).GetSupportedCurrencies(ctx, &v12.Empty{})
	if err1 != nil {
		fmt.Printf("could not charge the card: %+v", err)
		return
	}
	fmt.Println(resp1.GetCurrencyCodes())
}
