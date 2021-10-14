package biz

import "go.uber.org/zap"

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	PriceUsd    struct {
		CurrencyCode string
		Units        int64 `json:"units"`
		Nanos        int32 `json:"nanos"`
	} `json:"priceUsd"`
	Categories []string `json:"categories"`
}

type ProductCatalogUseCase struct {
	logger *zap.Logger
	data   []Product
}

func (c ProductCatalogUseCase) List(ctx interface{}) ([]Product, error) {
	return c.data, nil
}

func NewProductCatalogUseCase(logger *zap.Logger) *ProductCatalogUseCase {
	return &ProductCatalogUseCase{
		logger: logger,
	}
}
