package service

import (
	"order-inventory-transaction/model"
)

type OrderService interface {
	CreateOrder(request model.OrderRequest) (int, error)
	GetProducts() ([]model.Product, error)
}
