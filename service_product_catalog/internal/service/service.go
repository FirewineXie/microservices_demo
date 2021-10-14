package service

import (
	"go.uber.org/zap"
	v1 "microservices_demo_v1/service_product_catalog/api/v1"
	"microservices_demo_v1/service_product_catalog/internal/biz"

)

type ProductCatalogService struct {
	v1.UnimplementedProductCatalogServiceServer
	productCatalog *biz.ProductCatalogUseCase
	logger         *zap.Logger
}

func NewShippingService(productCatalog *biz.ProductCatalogUseCase, logger *zap.Logger) *ProductCatalogService {
	return &ProductCatalogService{
		productCatalog: productCatalog,
		logger:         logger.Named("service_product_catalog"),
	}
}