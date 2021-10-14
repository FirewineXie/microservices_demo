package service

import (
	"go.uber.org/zap"
	v1 "microservices_demo_v1/service_recommendation/api/v1"
	"microservices_demo_v1/service_recommendation/internal/biz"
)

type RecommendationService struct {
	v1.UnimplementedRecommendationServiceServer
	recommendation *biz.RecommendationUseCase
	logger         *zap.Logger
}

func NewRecommendationService(recommendation *biz.RecommendationUseCase, logger *zap.Logger) *RecommendationService {
	return &RecommendationService{
		recommendation: recommendation,
		logger:         logger.Named("service_recommendation"),
	}
}
