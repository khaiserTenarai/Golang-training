package service

import "context"

type OrderService interface {
	AddProduct(context.Context, string, float64, int) error
	CreateOrder(context.Context, int64, int) error
	ListProducts(context.Context) error
}
