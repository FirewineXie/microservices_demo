package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_recommendation/internal/api/v1"
	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx,
		"0.0.0.0:9008", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))

	recommendations, err := v1.NewRecommendationServiceClient(conn).ListRecommendations(ctx, &v1.ListRecommendationsRequest{})
	if err != nil {
		return
	}
	fmt.Println(recommendations.String())
}
