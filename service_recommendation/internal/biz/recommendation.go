package biz

import "go.uber.org/zap"

type RecommendationUseCase struct {
	logger *zap.Logger
}

func NewRecommendationUseCase(logger *zap.Logger) *RecommendationUseCase {
	return &RecommendationUseCase{
		logger: logger,
	}
}