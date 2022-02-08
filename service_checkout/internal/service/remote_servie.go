package service

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	v12 "microservices_demo/service_checkout/internal/api/v1"
)

func getUserCart(ctx context.Context, userID string) ([]*v12.CartItem, error) {

	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9002", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v12.NewCartServiceClient(conn).GetCart(ctx, &v12.GetCartRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("failed to get user cart during checkout: %+v", err)
	}
	return resp.GetItems(), nil
}

func quoteShipping(ctx context.Context, address *v12.Address, items []*v12.CartItem) (*v12.Money, error) {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9009", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	shippingQuote, err := v12.NewShippingServiceClient(conn).
		GetQuote(ctx, &v12.GetQuoteRequest{
			Address: address,
			Items:   items})
	if err != nil {
		return nil, fmt.Errorf("failed to get shipping quote: %+v", err)
	}
	return shippingQuote.GetCostUsd(), nil
}

func prepOrderItems(ctx context.Context, items []*v12.CartItem, userCurrency string) ([]*v12.OrderItem, error) {
	out := make([]*v12.OrderItem, len(items))

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9007", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cl := v12.NewProductCatalogServiceClient(conn)

	for i, item := range items {
		product, err := cl.GetProduct(ctx, &v12.GetProductRequest{Id: item.GetProductId()})
		if err != nil {
			return nil, fmt.Errorf("failed to get product #%q", item.GetProductId())
		}
		price, err := convertCurrency(ctx, product.GetPriceUsd(), userCurrency)
		if err != nil {
			return nil, fmt.Errorf("failed to convert price of %q to %s", item.GetProductId(), userCurrency)
		}
		out[i] = &v12.OrderItem{
			Item: item,
			Cost: price}
	}
	return out, nil
}
func convertCurrency(ctx context.Context, from *v12.Money, toCurrency string) (*v12.Money, error) {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9004", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := v12.NewCurrencyServiceClient(conn).Convert(context.TODO(), &v12.CurrencyConversionRequest{
		From:   from,
		ToCode: toCurrency})
	if err != nil {
		return nil, fmt.Errorf("failed to convert currency: %+v", err)
	}
	return result, err
}

func chargeCard(ctx context.Context, amount *v12.Money, paymentInfo *v12.CreditCardInfo) (string, error) {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9006", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	paymentResp, err := v12.NewPaymentServiceClient(conn).Charge(ctx, &v12.ChargeRequest{
		Amount:     amount,
		CreditCard: paymentInfo})
	if err != nil {
		return "", fmt.Errorf("could not charge the card: %+v", err)
	}
	return paymentResp.GetTransactionId(), nil
}

func sendOrderConfirmation(ctx context.Context, email string, order *v12.OrderResult) error {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9005", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = v12.NewEmailServiceClient(conn).SendOrderConfirmation(ctx, &v12.SendOrderConfirmationRequest{
		Email: email,
		Order: order})
	return err
}

func shipOrder(ctx context.Context, address *v12.Address, items []*v12.CartItem) (string, error) {

	opentracing.GlobalTracer()
	conn, err := grpc.DialContext(ctx,
		"0.0.0.0:9009", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	resp, err := v12.NewShippingServiceClient(conn).ShipOrder(ctx, &v12.ShipOrderRequest{
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

	if _, err = v12.NewCartServiceClient(conn).EmptyCart(ctx, &v12.EmptyCartRequest{UserId: userID}); err != nil {
		return fmt.Errorf("failed to empty user cart during checkout: %+v", err)
	}
	return nil
}
