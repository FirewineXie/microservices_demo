package service

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"math/rand"
	v1 "microservices_demo/gateway/api/v1"
	money "microservices_demo/gateway/internal/biz"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ctxKeyRequestID struct{}
type ctxKeySessionID struct{}

const (
	defaultCurrency = "USD"
	cookieMaxAge    = 60 * 60 * 48

	cookiePrefix    = "shop_"
	cookieSessionID = cookiePrefix + "session-id"
	cookieCurrency  = cookiePrefix + "currency"
)

func (gs *GatewayService) Router(ctx *gin.RouterGroup) {

	ctx.GET("/product/:id", gs.ProductHandler)
	ctx.GET("/cart", gs.ViewCartHandler)
	ctx.POST("/cart", gs.AddToCartHandler)
	ctx.POST("/cart/empty", gs.EmptyCartHandler)
	ctx.POST("/setCurrency", gs.SetCurrencyHandler)
	ctx.GET("/logout", gs.LogoutHandler)
	ctx.POST("/cart/checkout", gs.PlaceOrderHandler)

}

func (gs *GatewayService) HomeHandlerGet(ctx *gin.Context) {
	spanV, exists := ctx.Get("ctx")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": "could not retrieve continue",
		})
		return
	}
	spanCtx := spanV.(context.Context)
	currencies, err := getCurrencies(spanCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": "could not retrieve currency",
		})
		return
	}
	products, err := getProducts(spanCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": "could not retrieve product",
		})
		return
	}
	cart, err := getCart(spanCtx, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": "could not retrieve cart",
		})
		return
	}
	type productView struct {
		Item  *v1.Product
		Price *v1.Money
	}
	ps := make([]productView, len(products))
	for i, p := range products {
		price, err := convertCurrency(spanCtx, p.GetPriceUsd(), currentCurrency(ctx.Request))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"err": "could not retrieve currency",
			})
			return
		}
		ps[i] = productView{p, price}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"user_currency": currentCurrency(ctx.Request),
		"show_currency": true,
		"currencies":    currencies,
		"products":      ps,
		"cart_size":     cartSize(cart),
		"banner_color":  os.Getenv("BANNER_COLOR"), // illustrates canary deployments
		"ad":            chooseAd(spanCtx, []string{}),
	})
}

func (gs *GatewayService) ProductHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "product id not specified",
			"ok":  1,
		})
		return
	}

	p, err := getProduct(c.Request.Context(), id)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve product").Error(),
			"ok":  1,
		})
		return
	}
	currencies, err := getCurrencies(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve currencies"),
			"ok":  1,
		})
		return
	}
	cart, err := getCart(c.Request.Context(), sessionID(c.Request))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve product"),
			"ok":  1,
		})

		return
	}
	price, err := convertCurrency(c.Request.Context(),p.GetPriceUsd(),currentCurrency(c.Request))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve currencies"),
			"ok":  1,
		})
		return
	}
	recommendations, err := getRecommendations(c.Request.Context(), sessionID(c.Request), []string{id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": errors.Wrap(err, "failed to get product recommendations"),
			"ok":  1,
		})
		return
	}
	products := struct {
		Item  *v1.Product
		Price *v1.Money
	}{p, price}
	c.JSON(http.StatusOK, map[string]interface{}{
		"session_id":      sessionID(c.Request),
		"request_id":      c.Request.Context().Value(ctxKeyRequestID{}),
		"ad":              chooseAd(c.Request.Context(), p.Categories),
		"user_currency":   currentCurrency(c.Request),
		"show_currency":   true,
		"currencies":      currencies,
		"product":         products,
		"recommendations": recommendations,
		"cart_size":       cartSize(cart),
	})

}

func (gs *GatewayService) ViewCartHandler(c *gin.Context) {
	r := c.Request
	currencies, err := getCurrencies(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve currencies"),
			"ok":  1,
		})
		return
	}
	cart, err := getCart(r.Context(), sessionID(r))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve cart"),
			"ok":  1,
		})
		return
	}

	recommendations, err := getRecommendations(r.Context(), sessionID(r), cartIDs(cart))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "failed to get product recommendations"),
			"ok":  1,
		})
		return
	}

	shippingCost, err := getShippingQuote(r.Context(), cart, currentCurrency(r))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "failed to get shipping quote"),
			"ok":  1,
		})
		return
	}

	type cartItemView struct {
		Item     *v1.Product
		Quantity int32
		Price    *v1.Money
	}
	items := make([]cartItemView, len(cart))
	totalPrice := v1.Money{CurrencyCode: currentCurrency(r)}
	for i, item := range cart {
		p, err := getProduct(r.Context(), item.GetProductId())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"msg": errors.Wrapf(err, "could not retrieve product #%s", item.GetProductId()),
				"ok":  1,
			})
			return
		}
		price, err := convertCurrency(r.Context(), p.GetPriceUsd(), currentCurrency(r))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"msg": errors.Wrapf(err, "could not convert currency for product #%s", item.GetProductId()),
				"ok":  1,
			})
			return
		}

		multPrice := money.MultiplySlow(*price, uint32(item.GetQuantity()))
		items[i] = cartItemView{
			Item:     p,
			Quantity: item.GetQuantity(),
			Price:    &multPrice}
		totalPrice = money.Must(money.Sum(totalPrice, multPrice))
	}
	totalPrice = money.Must(money.Sum(totalPrice, *shippingCost))

	year := time.Now().Year()
	c.JSON(http.StatusOK, map[string]interface{}{
		"session_id":       sessionID(r),
		"request_id":       c.Request.Context().Value(ctxKeyRequestID{}),
		"user_currency":    currentCurrency(r),
		"currencies":       currencies,
		"recommendations":  recommendations,
		"cart_size":        cartSize(cart),
		"shipping_cost":    shippingCost,
		"show_currency":    true,
		"total_cost":       totalPrice,
		"items":            items,
		"expiration_years": []int{year, year + 1, year + 2, year + 3, year + 4},
	})

}

func (gs *GatewayService) AddToCartHandler(c *gin.Context) {
	r := c.Request
	quantity, _ := strconv.ParseUint(r.FormValue("quantity"), 10, 32)
	productID := r.FormValue("product_id")
	if productID == "" || quantity == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.New("invalid form input"),
			"ok":  1,
		})
		return
	}
	p, err := getProduct(r.Context(), productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve product"),
			"ok":  1,
		})

		return
	}

	if err := insertCart(r.Context(), sessionID(r), p.GetId(), int32(quantity)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "failed to add to cart"),
			"ok":  1,
		})

		return
	}
	c.Request.Response.Header.Set("location", "/cart")

	c.JSON(http.StatusFound, nil)
}

func (gs *GatewayService) EmptyCartHandler(c *gin.Context) {
	if err := emptyCart(c.Request.Context(), sessionID(c.Request)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "failed to empty cart"),
			"ok":  1,
		})
		return
	}
	c.Request.Response.Header.Set("location", "/")

	c.JSON(http.StatusFound, nil)

}

func (gs *GatewayService) SetCurrencyHandler(c *gin.Context) {
	cur := c.Request.FormValue("currency_code")

	if cur != "" {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   cookieCurrency,
			Value:  cur,
			MaxAge: cookieMaxAge,
		})
	}
	referer := c.Request.Header.Get("referer")
	if referer == "" {
		referer = "/"
	}

	c.Request.Response.Header.Set("location", referer)

	c.JSON(http.StatusFound, nil)
}

func (gs *GatewayService) LogoutHandler(c *gin.Context) {
	for _, cc := range c.Request.Cookies() {
		cc.Expires = time.Now().Add(-time.Hour * 24 * 365)
		cc.MaxAge = -1
		http.SetCookie(c.Writer, cc)
	}
	c.Request.Response.Header.Set("Location", "/")

	c.JSON(http.StatusFound, nil)
}

func (gs *GatewayService) PlaceOrderHandler(c *gin.Context) {
	r := c.Request
	var (
		email         = r.FormValue("email")
		streetAddress = r.FormValue("street_address")
		zipCode, _    = strconv.ParseInt(r.FormValue("zip_code"), 10, 32)
		city          = r.FormValue("city")
		state         = r.FormValue("state")
		country       = r.FormValue("country")
		ccNumber      = r.FormValue("credit_card_number")
		ccMonth, _    = strconv.ParseInt(r.FormValue("credit_card_expiration_month"), 10, 32)
		ccYear, _     = strconv.ParseInt(r.FormValue("credit_card_expiration_year"), 10, 32)
		ccCVV, _      = strconv.ParseInt(r.FormValue("credit_card_cvv"), 10, 32)
	)

	conn, err := grpc.DialContext(c.Request.Context(),
		serviceCheckout, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	order, err := v1.NewCheckoutServiceClient(conn).
		PlaceOrder(r.Context(), &v1.PlaceOrderRequest{
			Email: email,
			CreditCard: &v1.CreditCardInfo{
				CreditCardNumber:          ccNumber,
				CreditCardExpirationMonth: int32(ccMonth),
				CreditCardExpirationYear:  int32(ccYear),
				CreditCardCvv:             int32(ccCVV)},
			UserId:       sessionID(r),
			UserCurrency: currentCurrency(r),
			Address: &v1.Address{
				StreetAddress: streetAddress,
				City:          city,
				State:         state,
				ZipCode:       int32(zipCode),
				Country:       country},
		})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "failed to complete the order"),
			"ok":  1,
		})
		return
	}

	order.GetOrder().GetItems()
	recommendations, _ := getRecommendations(r.Context(), sessionID(r), nil)

	totalPaid := *order.GetOrder().GetShippingCost()
	for _, v := range order.GetOrder().GetItems() {
		totalPaid = money.Must(money.Sum(totalPaid, *v.GetCost()))
	}

	currencies, err := getCurrencies(r.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"msg": errors.Wrap(err, "could not retrieve currencies"),
			"ok":  1,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"session_id":      sessionID(r),
		"request_id":      r.Context().Value(ctxKeyRequestID{}),
		"user_currency":   currentCurrency(r),
		"show_currency":   false,
		"currencies":      currencies,
		"order":           order.GetOrder(),
		"total_paid":      &totalPaid,
		"recommendations": recommendations,
	})
}

// get total # of items in cart
func cartSize(c []*v1.CartItem) int {
	cartSize := 0
	for _, item := range c {
		cartSize += int(item.GetQuantity())
	}
	return cartSize
}

func chooseAd(ctx context.Context, ctxKeys []string) *v1.Ad {

	ads, err := getAd(ctx, ctxKeys)
	if err != nil {
		return nil
	}
	return ads[rand.Intn(len(ads))]
}

func currentCurrency(r *http.Request) string {
	c, _ := r.Cookie(cookieCurrency)
	if c != nil {
		return c.Value
	}
	return defaultCurrency
}
func sessionID(r *http.Request) string {
	v := r.Context().Value(ctxKeySessionID{})
	if v != nil {
		return v.(string)
	}
	return ""
}
func cartIDs(c []*v1.CartItem) []string {
	out := make([]string, len(c))
	for i, v := range c {
		out[i] = v.GetProductId()
	}
	return out
}
