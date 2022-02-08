package biz

import (
	"context"
	"go.uber.org/zap"
)

type Cart struct {
	Id int64

	ProductId string
	Quantity  int32
}

type CartUseCase struct {
	logger *zap.Logger
}

func (c CartUseCase) AddItem(ctx context.Context, c2 *Cart, userId string) error {
	c.logger.Info("AddItem called with ", zap.String("user_id", userId), zap.String("product_id", c2.ProductId), zap.Int32("quantity", c2.Quantity))
	userCartItems[userId] = append(userCartItems[userId], c2)
	return nil
}

func (c CartUseCase) ListCart(ctx context.Context, userId string) ([]*Cart, error) {
	if carts, ok := userCartItems[userId]; ok {

		return carts, nil
	}

	return nil, nil
}

func (c CartUseCase) DeleteCart(ctx context.Context, userId string) error {
	delete(userCartItems, userId)
	return nil
}

func NewCartUseCase(logger *zap.Logger) *CartUseCase {
	return &CartUseCase{logger}
}

var userCartItems map[string][]*Cart

func init() {
	userCartItems = make(map[string][]*Cart, 0)
}
