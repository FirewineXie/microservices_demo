package service

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	v1 "microservices_demo/gateway/api/v1"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	serviceAd             = "0.0.0.0:9001"
	serviceCart           = "0.0.0.0:9002"
	serviceCheckout       = "0.0.0.0:9003"
	serviceCurrency       = "0.0.0.0:9004"
	serviceEmail          = "0.0.0.0:9005"
	servicePayment        = "0.0.0.0:9006"
	serviceProductCatalog = "0.0.0.0:9007"
	serviceRecommendation = "0.0.0.0:9008"
	serviceShipping       = "0.0.0.0:9009"
)

func getProduct(ctx context.Context, id string) (*v1.Product, error) {
	conn, err := grpc.DialContext(ctx,
		serviceProductCatalog, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewProductCatalogServiceClient(conn).GetProduct(context.Background(), &v1.GetProductRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func getCurrencies(ctx context.Context) (*v1.GetSupportedCurrenciesResponse, error) {

	conn, err := grpc.DialContext(ctx,
		serviceCurrency, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewCurrencyServiceClient(conn).GetSupportedCurrencies(ctx, &v1.Empty{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getCart(ctx context.Context, userId string) ([]*v1.CartItem, error) {

	conn, err := grpc.DialContext(ctx,
		serviceCart, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewCartServiceClient(conn).GetCart(ctx, &v1.GetCartRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}
func emptyCart(ctx context.Context, userID string) error {

	conn, err := grpc.DialContext(ctx,
		serviceCart, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = v1.NewCartServiceClient(conn).EmptyCart(ctx, &v1.EmptyCartRequest{
		UserId: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func insertCart(ctx context.Context, userID, productID string, quantity int32) error {

	conn, err := grpc.DialContext(ctx,
		serviceCart, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = v1.NewCartServiceClient(conn).AddItem(ctx, &v1.AddItemRequest{
		UserId: userID,
		Item: &v1.CartItem{
			ProductId: productID,
			Quantity:  quantity},
	})
	return err
}

func convertCurrency(ctx context.Context, money *v1.Money, currency string) (*v1.Money, error) {

	conn, err := grpc.DialContext(ctx,
		serviceCurrency, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return v1.NewCurrencyServiceClient(conn).
		Convert(ctx, &v1.CurrencyConversionRequest{
			From:   money,
			ToCode: currency})
}

func getShippingQuote(ctx context.Context, items []*v1.CartItem, currency string) (*v1.Money, error) {

	conn, err := grpc.DialContext(ctx,
		serviceShipping, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	quote, err := v1.NewShippingServiceClient(conn).GetQuote(ctx,
		&v1.GetQuoteRequest{
			Address: nil,
			Items:   items})
	if err != nil {
		return nil, err
	}
	localized, err := convertCurrency(ctx, quote.GetCostUsd(), currency)
	return localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

func getRecommendations(ctx context.Context, userID string, productIDs []string) ([]*v1.Product, error) {

	conn, err := grpc.DialContext(ctx,
		serviceRecommendation, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewRecommendationServiceClient(conn).ListRecommendations(ctx,
		&v1.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Product, len(resp.GetProductIds()))
	for i, v := range resp.GetProductIds() {
		p, err := getProduct(ctx, v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}
	return out, err
}

func getAd(ctx context.Context, ctxKeys []string) ([]*v1.Ad, error) {

	conn, err := grpc.DialContext(ctx,
		serviceAd, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewAdServiceClient(conn).GetAds(ctx, &v1.AdRequest{
		ContextKeys: ctxKeys,
	})
	return resp.GetAds(), errors.Wrap(err, "failed to get ads")
}

func getProducts(ctx context.Context) ([]*v1.Product, error) {

	conn, err := grpc.DialContext(ctx,
		serviceProductCatalog, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := v1.NewProductCatalogServiceClient(conn).ListProducts(ctx, &v1.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
}
func sendOrderConfirmation(ctx context.Context, email string, order *v1.OrderResult) error {


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