package service

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	v1 "microservices_demo/service_checkout/api/v1"
)

func getUserCart(ctx context.Context, userID string) ([]*v1.CartItem, error) {

	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9002", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewCartServiceClient(conn).GetCart(ctx, &v1.GetCartRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("failed to get user cart during checkout: %+v", err)
	}
	return resp.GetItems(), nil
}

func quoteShipping(ctx context.Context, address *v1.Address, items []*v1.CartItem) (*v1.Money, error) {
	

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9009", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	shippingQuote, err := v1.NewShippingServiceClient(conn).
		GetQuote(ctx, &v1.GetQuoteRequest{
			Address: address,
			Items:   items})
	if err != nil {
		return nil, fmt.Errorf("failed to get shipping quote: %+v", err)
	}
	return shippingQuote.GetCostUsd(), nil
}

func prepOrderItems(ctx context.Context, items []*v1.CartItem, userCurrency string) ([]*v1.OrderItem, error) {
	out := make([]*v1.OrderItem, len(items))
	
	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9007", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cl := v1.NewProductCatalogServiceClient(conn)

	for i, item := range items {
		product, err := cl.GetProduct(ctx, &v1.GetProductRequest{Id: item.GetProductId()})
		if err != nil {
			return nil, fmt.Errorf("failed to get product #%q", item.GetProductId())
		}
		price, err := convertCurrency(ctx, product.GetPriceUsd(), userCurrency)
		if err != nil {
			return nil, fmt.Errorf("failed to convert price of %q to %s", item.GetProductId(), userCurrency)
		}
		out[i] = &v1.OrderItem{
			Item: item,
			Cost: price}
	}
	return out, nil
}
func convertCurrency(ctx context.Context, from *v1.Money, toCurrency string) (*v1.Money, error) {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9004", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := v1.NewCurrencyServiceClient(conn).Convert(context.TODO(), &v1.CurrencyConversionRequest{
		From:   from,
		ToCode: toCurrency})
	if err != nil {
		return nil, fmt.Errorf("failed to convert currency: %+v", err)
	}
	return result, err
}

func chargeCard(ctx context.Context, amount *v1.Money, paymentInfo *v1.CreditCardInfo) (string, error) {


	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9006", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	paymentResp, err := v1.NewPaymentServiceClient(conn).Charge(ctx, &v1.ChargeRequest{
		Amount:     amount,
		CreditCard: paymentInfo})
	if err != nil {
		return "", fmt.Errorf("could not charge the card: %+v", err)
	}
	return paymentResp.GetTransactionId(), nil
}

func sendOrderConfirmation(ctx context.Context, email string, order *v1.OrderResult) error {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9005", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = v1.NewEmailServiceClient(conn).SendOrderConfirmation(ctx, &v1.SendOrderConfirmationRequest{
		Email: email,
		Order: order})
	return err
}

func shipOrder(ctx context.Context, address *v1.Address, items []*v1.CartItem) (string, error) {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9009", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	resp, err := v1.NewShippingServiceClient(conn).ShipOrder(ctx, &v1.ShipOrderRequest{
		Address: address,
		Items:   items})
	if err != nil {
		return "", fmt.Errorf("shipment failed: %+v", err)
	}
	return resp.GetTrackingId(), nil
}

func emptyUserCart(ctx context.Context, userID string) error {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9002", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err = v1.NewCartServiceClient(conn).EmptyCart(ctx, &v1.EmptyCartRequest{UserId: userID}); err != nil {
		return fmt.Errorf("failed to empty user cart during checkout: %+v", err)
	}
	return nil
}
