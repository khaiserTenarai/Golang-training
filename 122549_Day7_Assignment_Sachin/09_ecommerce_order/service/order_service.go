package service

import "context"

type OrderService interface {
	AddCustomer(context.Context, string, string) error
	AddProduct(context.Context, string, float64) error
	CreateOrder(context.Context, int64, int64, int) error
	ShowOrders(context.Context) error
}
