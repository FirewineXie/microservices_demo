package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "microservices_demo/service_cart/api/v1"
	"microservices_demo/service_cart/internal/biz"
)

func (cs *CartService) AddItem(c context.Context, request *v1.AddItemRequest) (*v1.Empty, error) {
	item := request.GetItem()
	cart := biz.Cart{
		ProductId: item.GetProductId(),
		Quantity:  item.GetQuantity(),
	}
	err := cs.cart.AddItem(c, &cart, request.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "无法进行数据")
	}
	return nil, nil
}
func (cs *CartService) GetCart(ctx context.Context, request *v1.GetCartRequest) (*v1.Cart, error) {
	listCart, err := cs.cart.ListCart(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "not get cart data ")
	}
	var cartItem []*v1.CartItem
	for _, cart := range listCart {
		cartItem = append(cartItem, &v1.CartItem{
			ProductId: cart.ProductId,
			Quantity:  cart.Quantity,
		})
	}
	result := v1.Cart{
		UserId: request.GetUserId(),
		Items:  cartItem,
	}
	return &result, nil
}
func (cs *CartService) EmptyCart(ctx context.Context, request *v1.EmptyCartRequest) (*v1.Empty, error) {
	err := cs.cart.DeleteCart(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete error")
	}
	return nil, nil
}
