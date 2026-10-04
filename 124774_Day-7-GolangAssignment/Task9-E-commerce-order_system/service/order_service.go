package service

import "ecommerce-order/model"

type OrderService interface {
	CreateCustomer(customer model.Customer) error

	CreateProduct(product model.Product) error

	CreateOrder(customerID int) (int, error)

	AddOrderItem(item model.OrderItem) error

	GetOrderDetails(orderID int) ([]model.OrderDetail, error)
}
