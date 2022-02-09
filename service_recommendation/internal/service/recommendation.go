package service

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "microservices_demo/service_recommendation/api/v1"
)

func (rs *RecommendationService) ListRecommendations(c context.Context, request *v1.ListRecommendationsRequest) (
	*v1.ListRecommendationsResponse, error) {
	//maxResponse := 5
	products, err := getProducts(c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error")
	}
	var productIds []string
	productIdsMap := map[string]bool{}
	for _, x := range products {
		productIds = append(productIds, x.Id)
		productIdsMap[x.Id] = true
	}
	for _, x := range request.ProductIds {
		delete(productIdsMap, x)
	}
	var filteredProducts []string
	for k, _ := range productIdsMap {
		filteredProducts = append(filteredProducts, k)
	}

	rs.logger.Info("resv listRecommendations product_ids :", zap.Any("productIds", filteredProducts))

	response := v1.ListRecommendationsResponse{}
	response.ProductIds = append(response.ProductIds, filteredProducts...)

	return &response,nil
}
