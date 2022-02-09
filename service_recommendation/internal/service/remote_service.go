package service

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v12 "microservices_demo/service_recommendation/internal/api/v1"
)

var (
	serviceProductCatalog = "0.0.0.0:9007"
)

func getProducts(ctx context.Context) ([]*v12.Product, error) {
	var response []*v12.Product
	conn, err := grpc.DialContext(ctx,
		serviceProductCatalog, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return response, err
	}
	defer conn.Close()
	resp, err := v12.NewProductCatalogServiceClient(conn).ListProducts(ctx, &v12.Empty{})
	if err != nil {
		return response, err
	}
	return resp.Products, nil
}
