package service

import (
	"context"
	"errors"
	"example.com/q9-ecommerce-order/repository"
)

type OrderServiceImpl struct{ repo repository.OrderRepository }

func NewOrderService(r repository.OrderRepository) OrderService { return &OrderServiceImpl{repo: r} }
func (s *OrderServiceImpl) AddCustomer(c context.Context, n, e string) error {
	if n == "" || e == "" {
		return errors.New("name and email are required")
	}
	return s.repo.AddCustomer(c, n, e)
}
func (s *OrderServiceImpl) AddProduct(c context.Context, n string, p float64) error {
	if n == "" || p <= 0 {
		return errors.New("product name and positive price are required")
	}
	return s.repo.AddProduct(c, n, p)
}
func (s *OrderServiceImpl) CreateOrder(c context.Context, cid, pid int64, q int) error {
	return s.repo.CreateOrder(c, cid, pid, q)
}
func (s *OrderServiceImpl) ShowOrders(c context.Context) error { return s.repo.ShowOrders(c) }
