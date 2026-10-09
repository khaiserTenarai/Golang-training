package repository

import "order-inventory-transaction/model"

type OrderRepository interface {
	CreateOrder(order model.Order, items []model.OrderItem) (int, error)
	GetProducts() ([]model.Product, error)
	GetProduct(productID int) (model.Product, error)
}
