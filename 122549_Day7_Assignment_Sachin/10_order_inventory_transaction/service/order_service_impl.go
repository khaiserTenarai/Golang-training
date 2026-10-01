package service

import (
	"context"
	"errors"
	"example.com/q10-order-inventory/repository"
)

type OrderServiceImpl struct{ repo repository.OrderRepository }

func NewOrderService(r repository.OrderRepository) OrderService { return &OrderServiceImpl{repo: r} }
func (s *OrderServiceImpl) AddProduct(c context.Context, n string, p float64, q int) error {
	if n == "" || p <= 0 || q < 0 {
		return errors.New("invalid product details")
	}
	return s.repo.AddProduct(c, n, p, q)
}
func (s *OrderServiceImpl) CreateOrder(c context.Context, p int64, q int) error {
	return s.repo.CreateOrder(c, p, q)
}
func (s *OrderServiceImpl) ListProducts(c context.Context) error { return s.repo.ListProducts(c) }
