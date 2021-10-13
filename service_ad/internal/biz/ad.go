package biz

import "errors"

type Ad struct {
	RedirectUrl string
	Text        string
}




var adsMap map[string][]Ad

func init() {
	adsMap = make(map[string][]Ad)
	camera := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	lens := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	recordPlayer := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	bike := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	baristaKit := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	airPlant := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	terrarium := Ad{
		RedirectUrl: "/product/2ZYFJ3GM2N",
		Text: "Film camera for sale 。 50% off.",
	}
	photography := []Ad{camera,lens}
	vintage := []Ad{camera, lens, recordPlayer}
	cycling := []Ad{bike}
	cookware := []Ad{baristaKit}
	gardening := []Ad{airPlant, terrarium}

	adsMap["photography"] = photography
	adsMap["vintage"] = vintage
	adsMap["cycling"] = cycling
	adsMap["cookware"] = cookware
	adsMap["gardening"] = gardening
}
func  GetAllAdsList() (ads [][]Ad, err error) {
	for _, v := range adsMap {
		ads = append(ads, v)
	}
	return
}

func GetAdsByCategory(category string) (ads []Ad, err error) {

	if adList, ok := adsMap[category]; ok {
		return adList, nil
	}
	return nil, errors.New("not found")
}

