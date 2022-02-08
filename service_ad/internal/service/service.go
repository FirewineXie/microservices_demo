package service

import (
	"go.uber.org/zap"
	"math/rand"
	v1 "microservices_demo/service_ad/internal/api/v1"
	"microservices_demo/service_ad/internal/repo"

	"microservices_demo/service_ad/internal/biz"
)

type AdService struct {
	v1.UnimplementedAdServiceServer
	logger *zap.Logger
}

func NewAdService(logger *zap.Logger) *AdService {
	return &AdService{
		logger: logger,
	}
}

const (
	MAX_ADS_TO_SERVE = 2
)

func GetAdsByCategory(category string) ([]biz.Ad, error) {
	return repo.GetAdsByCategory(category)
}

func GetRandomAds() []biz.Ad {
	var ads []biz.Ad
	allAds, _ := repo.GetAllAdsList()
	for i := 0; i < MAX_ADS_TO_SERVE; i++ {
		n := rand.Intn(len(allAds))
		ads = append(ads, allAds[n]...)
	}
	return ads
}
