package service

import (
	"context"
	"math"
	v12 "microservices_demo/service_currency/internal/api/v1"
	"microservices_demo/service_currency/internal/biz"
)

// GetSupportedCurrencies Lists the supported currencies
func (cs *CurrencyService) GetSupportedCurrencies(context.Context, *v12.Empty) (*v12.GetSupportedCurrenciesResponse, error) {
	cs.logger.Info("Getting supported currencies...")
	currencyData := cs.shipping.GetCurrencyData()
	var keys []string
	for k, _ := range currencyData {
		keys = append(keys, k)
	}

	return &v12.GetSupportedCurrenciesResponse{
		CurrencyCodes: keys,
	}, nil
}
func (cs *CurrencyService) Convert(c context.Context, request *v12.CurrencyConversionRequest) (*v12.Money, error) {
	cs.logger.Info("received conversion request")

	from := request.From

	data := cs.shipping.GetCurrencyData()

	euros := biz.Carry(biz.Money{
		Units: int64(float32(from.Units) / data[from.CurrencyCode]),
		Nanos: int32(float32(from.Nanos) / data[from.CurrencyCode]),
	})

	euros.Nanos = int32(math.Round(float64(euros.Nanos)))
	// Convert: EUR --> to_currency
	result := biz.Carry(biz.Money{
		Units: int64(float32(from.Units) / data[request.ToCode]),
		Nanos: int32(float32(from.Nanos) / data[request.ToCode]),
	})

	result.Units = int64(math.Floor(float64(result.Units)))
	result.Nanos = int32(math.Floor(float64(result.Nanos)))
	result.CurrencyCode = request.ToCode

	return &v12.Money{
		Units:        result.Units,
		Nanos:        result.Nanos,
		CurrencyCode: result.CurrencyCode,
	}, nil
}
