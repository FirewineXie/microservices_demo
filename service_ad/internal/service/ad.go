package service

import (
	"context"
	v1 "microservices_demo/service_ad/internal/api/v1"

	"microservices_demo/service_ad/internal/biz"
)

func (s *AdService) GetAds(ctx context.Context,
	req *v1.AdRequest) (resp *v1.AdResponse, err error) {

	resp = new(v1.AdResponse)

	var allAds []biz.Ad

	if len(req.GetContextKeys()) > 0 {
		for _, key := range req.ContextKeys {
			ads, _ := GetAdsByCategory(key)
			allAds = append(allAds, ads...)
		}
	} else {
		allAds = GetRandomAds()
	}
	if len(allAds) == 0 {
		allAds = GetRandomAds()
	}
	for _, ad := range allAds {
		resp.Ads = append(resp.Ads, &v1.Ad{
			RedirectUrl: ad.RedirectUrl,
			Text:        ad.Text,
		})
	}

	return resp, nil
}
