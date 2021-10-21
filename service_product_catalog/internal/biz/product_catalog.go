package biz

import (
	"encoding/json"
	"go.uber.org/zap"
)

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
	cc := ` [
        {
            "id": "OLJCESPC7Z",
            "name": "Vintage Typewriter",
            "description": "This typewriter looks good in your living room.",
            "picture": "/static/img/products/typewriter.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 67,
                "nanos": 990000000
            },
            "categories": ["vintage"]
        },
        {
            "id": "66VCHSJNUP",
            "name": "Vintage Camera Lens",
            "description": "You won't have a camera to use it and it probably doesn't work anyway.",
            "picture": "/static/img/products/camera-lens.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 12,
                "nanos": 490000000
            },
            "categories": ["photography", "vintage"]
        },
        {
            "id": "1YMWWN1N4O",
            "name": "Home Barista Kit",
            "description": "Always wanted to brew coffee with Chemex and Aeropress at home?",
            "picture": "/static/img/products/barista-kit.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 124
            },
            "categories": ["cookware"]
        },
        {
            "id": "L9ECAV7KIM",
            "name": "Terrarium",
            "description": "This terrarium will looks great in your white painted living room.",
            "picture": "/static/img/products/terrarium.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 36,
                "nanos": 450000000
            },
            "categories": ["gardening"]
        },
        {
            "id": "2ZYFJ3GM2N",
            "name": "Film Camera",
            "description": "This camera looks like it's a film camera, but it's actually digital.",
            "picture": "/static/img/products/film-camera.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 2245
            },
            "categories": ["photography", "vintage"]
        },
        {
            "id": "0PUK6V6EV0",
            "name": "Vintage Record Player",
            "description": "It still works.",
            "picture": "/static/img/products/record-player.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 65,
                "nanos": 500000000
            },
            "categories": ["music", "vintage"]
        },
        {
            "id": "LS4PSXUNUM",
            "name": "Metal Camping Mug",
            "description": "You probably don't go camping that often but this is better than plastic cups.",
            "picture": "/static/img/products/camp-mug.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 24,
                "nanos": 330000000
            },
            "categories": ["cookware"]
        },
        {
            "id": "9SIQT8TOJO",
            "name": "City Bike",
            "description": "This single gear bike probably cannot climb the hills of San Francisco.",
            "picture": "/static/img/products/city-bike.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 789,
                "nanos": 500000000
            },
            "categories": ["cycling"]
        },
        {
            "id": "6E92ZMYYFZ",
            "name": "Air Plant",
            "description": "Have you ever wondered whether air plants need water? Buy one and figure out.",
            "picture": "/static/img/products/air-plant.jpg",
            "priceUsd": {
                "currencyCode": "USD",
                "units": 12,
                "nanos": 300000000
            },
            "categories": ["gardening"]
        }
    ]
`
	var data []Product
	json.Unmarshal([]byte(cc),&data)
	return &ProductCatalogUseCase{
		data: data,
		logger: logger,
	}
}
