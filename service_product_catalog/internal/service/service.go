package service

import (
	"go.uber.org/zap"
	"microservices_demo/service_product_catalog/internal/api/v1"
	"microservices_demo/service_product_catalog/internal/biz"
)

type ProductCatalogService struct {
	v1.UnimplementedProductCatalogServiceServer
	productCatalog *biz.ProductCatalogUseCase
	logger         *zap.Logger
}

func NewProductCatalogService(productCatalog *biz.ProductCatalogUseCase, logger *zap.Logger) *ProductCatalogService {
	return &ProductCatalogService{
		productCatalog: productCatalog,
		logger:         logger.Named("service_product_catalog"),
	}
}
