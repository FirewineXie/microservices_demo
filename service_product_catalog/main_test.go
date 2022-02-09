package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"microservices_demo/service_product_catalog/internal/api/v1"

	"testing"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx,
		"0.0.0.0:9007", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))

	product, err := v1.NewProductCatalogServiceClient(conn).GetProduct(ctx, &v1.GetProductRequest{
		Id: "OLJCESPC7Z",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(product.String())
}
