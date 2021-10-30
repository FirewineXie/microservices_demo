package service

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_recommendation/api/v1"
)
var (
	serviceProductCatalog = "0.0.0.0:9007"
)
func getProducts(ctx context.Context) ([]*v1.Product, error) {

	conn, err := grpc.DialContext(ctx,
		serviceProductCatalog, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewProductCatalogServiceClient(conn).ListProducts(ctx, &v1.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
}