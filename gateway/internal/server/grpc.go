package server

//// NewGRPCServer new a gRPC server.
//func NewGRPCServer(logger *zap.Logger, server *service.ShippingService) *grpc.Server {
//
//	srv := grpc.NewServer(
//		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
//			grpc_prometheus.UnaryServerInterceptor,
//			grpc_validator.UnaryServerInterceptor(),
//			grpc_zap.UnaryServerInterceptor(logger),
//			grpc_recovery.UnaryServerInterceptor(),
//			grpc_opentracing.UnaryServerInterceptor(),
//		)))
//
//	v1.RegisterShippingServiceServer(srv,server)
//	return srv
//}
