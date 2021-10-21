package service

import (
	"context"
	"math"
	v1 "microservices_demo/service_currency/api/v1"
	"microservices_demo/service_currency/internal/biz"
)

// GetSupportedCurrencies Lists the supported currencies
func (cs *CurrencyService) GetSupportedCurrencies(context.Context, *v1.Empty) (*v1.GetSupportedCurrenciesResponse, error) {
	cs.logger.Info("Getting supported currencies...")
	currencyData := cs.shipping.GetCurrencyData()
	var keys []string
	for k,_ := range currencyData{
		keys = append(keys, k)
	}

	return &v1.GetSupportedCurrenciesResponse{
		CurrencyCodes: keys,
	},nil
}
func (cs *CurrencyService) Convert(c context.Context,request *v1.CurrencyConversionRequest) (*v1.Money, error) {
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

	return &v1.Money{
		Units:        result.Units,
		Nanos:        result.Nanos,
		CurrencyCode: result.CurrencyCode,
	}, nil
}