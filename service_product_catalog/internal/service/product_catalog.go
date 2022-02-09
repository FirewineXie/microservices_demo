package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "microservices_demo/service_product_catalog/api/v1"
	"strings"
)

func (pcs *ProductCatalogService) ListProducts(ctx context.Context,req *v1.Empty) (*v1.ListProductsResponse, error) {
	var products []*v1.Product
	catalog, _ := pcs.productCatalog.List(ctx)
	for _, product := range catalog {
		money := v1.Money{
			Nanos:        product.PriceUsd.Nanos,
			Units:        product.PriceUsd.Units,
			CurrencyCode: product.PriceUsd.CurrencyCode,
		}
		pt := v1.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
		}
		pt.PriceUsd = &money
		products = append(products, &pt)
	}
	return &v1.ListProductsResponse{
		Products: products,
	}, nil
}
func (pcs *ProductCatalogService)  GetProduct(ctx context.Context, req *v1.GetProductRequest) (*v1.Product, error) {
	var found *v1.Product
	catalog, err := pcs.productCatalog.List(ctx)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(catalog); i++ {
		if req.GetId() == catalog[i].Id {
			temp := v1.Product{}
			temp.Id = catalog[i].Id
			temp.Name = catalog[i].Name
			temp.Description = catalog[i].Description
			temp.Picture = catalog[i].Picture
			temp.PriceUsd = &v1.Money{
				Nanos:        catalog[i].PriceUsd.Nanos,
				Units:        catalog[i].PriceUsd.Units,
				CurrencyCode: catalog[i].PriceUsd.CurrencyCode,
			}
			found = &temp
			break
		}
	}
	if found == nil {
		return nil, status.Errorf(codes.NotFound, "no product with ID %s", req.Id)
	}
	return found, nil
}
func (pcs *ProductCatalogService)  SearchProducts(ctx context.Context, request *v1.SearchProductsRequest) (*v1.SearchProductsResponse, error) {
	var ps []*v1.Product
	catalog, err := pcs.productCatalog.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range catalog {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(request.Query)) ||
			strings.Contains(strings.ToLower(product.Description), strings.ToLower(request.Query)) {

			ps = append(ps, &v1.Product{
				Id:          product.Id,
				Name:        product.Name,
				Description: product.Description,
				Picture:     product.Picture,
				PriceUsd: &v1.Money{
					Nanos:        product.PriceUsd.Nanos,
					Units:        product.PriceUsd.Units,
					CurrencyCode: product.PriceUsd.CurrencyCode,
				},
			})
		}
	}
	return &v1.SearchProductsResponse{Results: ps}, nil
}