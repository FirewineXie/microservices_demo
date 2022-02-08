package repo

import (
	"errors"
	"microservices_demo/service_ad/internal/biz"
)

//
var adsMap map[string][]biz.Ad

func init() {
	adsMap = make(map[string][]biz.Ad)
	camera := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	lens := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	recordPlayer := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	bike := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	baristaKit := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	airPlant := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	terrarium := biz.Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text:        "Film camera for sale 。 50% off.",
	}
	photography := []biz.Ad{camera, lens}
	vintage := []biz.Ad{camera, lens, recordPlayer}
	cycling := []biz.Ad{bike}
	cookware := []biz.Ad{baristaKit}
	gardening := []biz.Ad{airPlant, terrarium}

	adsMap["photography"] = photography
	adsMap["vintage"] = vintage
	adsMap["cycling"] = cycling
	adsMap["cookware"] = cookware
	adsMap["gardening"] = gardening
}
func GetAllAdsList() (ads [][]biz.Ad, err error) {
	for _, v := range adsMap {
		ads = append(ads, v)
	}
	return
}

func GetAdsByCategory(category string) (ads []biz.Ad, err error) {

	if adList, ok := adsMap[category]; ok {
		return adList, nil
	}
	return nil, errors.New("not found")
}
