package biz

import (
	"encoding/json"
	"go.uber.org/zap"
)

type CurrencyUseCase struct {
	logger *zap.Logger
	data   map[string]float32
}

func (c CurrencyUseCase) GetCurrencyData() map[string]float32 {
	return c.data
}

func NewCurrencyUseCase(logger *zap.Logger) *CurrencyUseCase {
	var jd map[string]float32
	json.Unmarshal([]byte(Cc), &jd)

	return &CurrencyUseCase{
		logger: logger,
		data:   jd,
	}
}

const Cc = `{
  "EUR":1.0,
  "USD":1.1305,
  "JPY":126.40,
  "BGN":1.9558,
  "CZK":25.592,
  "DKK":7.4609,
  "GBP":0.85970,
  "HUF":315.51,
  "PLN":4.2996,
  "RON":4.7463,
  "SEK":10.5375,
  "CHF":1.1360,
  "ISK":136.80,
  "NOK":9.8040,
  "HRK":7.4210,
  "RUB":74.4208,
  "TRY":6.1247,
  "AUD":1.6072,
  "BRL":4.2682,
  "CAD":1.5128,
  "CNY":7.5857,
  "HKD":8.8743,
  "IDR":15999.40,
  "ILS":4.0875,
  "INR":79.4320,
  "KRW":1275.05,
  "MXN":21.7999,
  "MYR":4.6289,
  "NZD":1.6679,
  "PHP":59.083,
  "SGD":1.5349,
  "THB":36.012,
  "ZAR":16.0583
}`

