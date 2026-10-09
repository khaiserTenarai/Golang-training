package service

import (
	"errors"
	"strings"

	"order-inventory-transaction/model"

	"order-inventory-transaction/repository"
)

type OrderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(
	repository repository.OrderRepository,
) *OrderServiceImpl {
	return &OrderServiceImpl{
		repository: repository,
	}
}

func (s *OrderServiceImpl) CreateOrder(
	request model.OrderRequest,
) (int, error) {
	if strings.TrimSpace(request.CustomerName) == "" {
		return 0, errors.New("customer name cannot be empty")
	}

	if len(request.Items) == 0 {
		return 0, errors.New("order must contain at least one item")
	}

	for _, item := range request.Items {
		if item.ProductID <= 0 {
			return 0, errors.New("product ID must be greater than zero")
		}

		if item.Quantity <= 0 {
			return 0, errors.New("quantity must be greater than zero")
		}
	}

	order := model.Order{
		CustomerName: request.CustomerName,
		Status:       "PLACED",
	}

	return s.repository.CreateOrder(order, request.Items)
}

func (s *OrderServiceImpl) GetProducts() ([]model.Product, error) {
	return s.repository.GetProducts()
}
