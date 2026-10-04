package repository

import "ecommerce-order/model"

type OrderRepository interface {
	CreateCustomer(customer model.Customer) error

	CreateProduct(product model.Product) error

	CreateOrder(customerID int) (int, error)

	CreateOrderItem(item model.OrderItem) error

	GetOrderDetails(orderID int) ([]model.OrderDetail, error)
}
