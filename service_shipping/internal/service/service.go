package service

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	v1 "microservices_demo/service_shipping/api/v1"
)

func (ss *ShippingService) GetQuote(c context.Context, request *v1.GetQuoteRequest) (*v1.GetQuoteResponse, error) {

	ss.logger.Info("[GetQuote] received request")
	defer ss.logger.Info("[GetQuote completed request")

	count := 0
	for _, item := range request.Items {
		count += int(item.Quantity)
	}

	quote := ss.shipping.CreateQuoteFromCount(count)

	return &v1.GetQuoteResponse{
		CostUsd: &v1.Money{
			CurrencyCode: "USD",
			Units:        int64(quote.Dollars),
			Nanos:        int32(quote.Cents * 10000000)},
	}, nil

}
func (ss *ShippingService) ShipOrder(c context.Context, request *v1.ShipOrderRequest) (*v1.ShipOrderResponse, error) {
	log.Info("[ShipOrder] received request")
	defer log.Info("[ShipOrder] completed request")
	// 1. Create a Tracking ID
	baseAddress := fmt.Sprintf("%s, %s, %s", request.Address.StreetAddress, request.Address.City, request.Address.State)
	id := ss.shipping.CreateTrackingId(baseAddress)

	// 2. Generate a response.
	return &v1.ShipOrderResponse{
		TrackingId: id,
	}, nil
}
