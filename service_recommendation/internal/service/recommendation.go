package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "microservices_demo_v1/service_recommendation/api/v1"
)

func (rs *RecommendationService) ListRecommendations(c context.Context, request *v1.ListRecommendationsRequest) (
	*v1.ListRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecommendations not implemented")
}
