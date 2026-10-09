package repository

import "ecommerce-order-system/model"

type EcommerceRepository interface {
	CreateCustomer(customer model.Customer) error
	CreateProduct(product model.Product) error
	CreateOrder(order model.Order) (int, error)
	AddOrderItem(item model.OrderItem) error

	GetCustomers() ([]model.Customer, error)
	GetProducts() ([]model.Product, error)
	GetOrderDetails(orderID int) ([]model.OrderDetails, error)
	GetAllOrderDetails() ([]model.OrderDetails, error)
}
