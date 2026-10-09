package service

import "ecommerce-order-system/model"

type EcommerceService interface {
	CreateCustomer(customer model.Customer) error
	CreateProduct(product model.Product) error
	CreateOrder(order model.Order, items []model.OrderItem) (int, error)

	GetCustomers() ([]model.Customer, error)
	GetProducts() ([]model.Product, error)
	GetOrderDetails(orderID int) ([]model.OrderDetails, error)
	GetAllOrderDetails() ([]model.OrderDetails, error)
}
