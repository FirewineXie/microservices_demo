package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_ad/internal/api/v1"

	"microservices_demo/service_ad/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(logger *zap.Logger, blog *service.AdService) *grpc.Server {

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,    // Prometheus
			grpc_validator.UnaryServerInterceptor(),   // 参数校验
			grpc_zap.UnaryServerInterceptor(logger),   // 日志流输出
			grpc_recovery.UnaryServerInterceptor(),    // recovery
			grpc_opentracing.UnaryServerInterceptor(), // 链路追踪
		)))
	v1.RegisterAdServiceServer(srv, blog)
	return srv
}
