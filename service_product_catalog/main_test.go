package main

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_product_catalog/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx,
		"0.0.0.0:9000", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))

	v1.NewProductCatalogServiceClient(conn).GetProduct(ctx,&v1.GetProductRequest{})
}
