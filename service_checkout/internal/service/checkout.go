package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	v1 "microservices_demo/service_checkout/api/v1"
)

func (cs  *CheckoutService) PlaceOrder(ctx context.Context,request *v1.PlaceOrderRequest) (resp *v1.PlaceOrderResponse,err error) {

	orderID, err := uuid.NewUUID()
	//
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to generate order uuid")
	//}
	//
	//prep, err := cs.prepareOrderItemsAndShippingQuoteFromCart(ctx, request.UserId, request.UserCurrency, request.Address)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, err.Error())
	//}
	//total := biz.Money{
	//	CurrencyCode: request.UserCurrency,
	//	Units:        0,
	//	Nanos:        0}
	//
	//total = biz.Must(biz.Sum(total, biz.Money{}))
	//for _, it := range prep.orderItems {
	//	total = biz.Must(biz.Sum(total, biz.Money{
	//		it.Cost.CurrencyCode,
	//		it.Cost.Units,
	//		it.Cost.Nanos,
	//	}))
	//}
	//
	//txID, err := chargeCard(ctx, &v1.Money{
	//	Nanos: total.Nanos,
	//	CurrencyCode: total.CurrencyCode,
	//	Units: total.Units,
	//}, request.CreditCard)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to charge card: %+v", err)
	//}
	//
	//cs.logger.Info("payment went through ", zap.String("transaction_id", txID))
	//shippingTrackingID, err := shipOrder(ctx, request.Address, prep.cartItems)
	//if err != nil {
	//	return nil, status.Errorf(codes.Unavailable, "shipping error: %+v", err)
	//}
	//
	//_ = emptyUserCart(ctx, request.UserId)
	//
	orderResult := &v1.OrderResult{
		OrderId:            orderID.String(),
		//ShippingTrackingId: shippingTrackingID,
		//ShippingCost:       prep.shippingCostLocalized,
		ShippingAddress:    request.Address,
		//Items:              prep.orderItems,
	}

	if err := sendOrderConfirmation(ctx, request.Email, orderResult); err != nil {
		cs.logger.Warn("failed to send order confirmation to", zap.String("email", request.Email), zap.Error(err))
	} else {
		cs.logger.Info("order confirmation email sent to", zap.String("email", request.Email))
	}
	resp = &v1.PlaceOrderResponse{Order: orderResult}
	return resp, nil
}


type orderPrep struct {
	orderItems            []*v1.OrderItem
	cartItems             []*v1.CartItem
	shippingCostLocalized *v1.Money
}

func (s *CheckoutService) prepareOrderItemsAndShippingQuoteFromCart(ctx context.Context, userID, userCurrency string, address *v1.Address) (
	orderPrep, error) {
	var out orderPrep
	cartItems, err := getUserCart(ctx, userID)
	if err != nil {
		return out, fmt.Errorf("cart failure: %+v", err)
	}
	orderItems, err := prepOrderItems(ctx, cartItems, userCurrency)
	if err != nil {
		return out, fmt.Errorf("failed to prepare order: %+v", err)
	}
	shippingUSD, err := quoteShipping(ctx, address, cartItems)
	if err != nil {
		return out, fmt.Errorf("shipping quote failure: %+v", err)
	}
	shippingPrice, err := convertCurrency(ctx, shippingUSD, userCurrency)
	if err != nil {
		return out, fmt.Errorf("failed to convert shipping cost to currency: %+v", err)
	}

	out.shippingCostLocalized = shippingPrice
	out.cartItems = cartItems
	out.orderItems = orderItems
	return out, nil
}
