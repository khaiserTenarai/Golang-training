package service

import (
	"errors"

	"ecommerce-order/model"
	"ecommerce-order/repository"
	"ecommerce-order/utility"
)

type OrderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(
	repository repository.OrderRepository,
) OrderService {

	return &OrderServiceImpl{
		repository: repository,
	}
}

func (s *OrderServiceImpl) CreateCustomer(
	customer model.Customer,
) error {

	if err := utility.ValidateName(customer.Name); err != nil {
		return err
	}

	if err := utility.ValidateEmail(customer.Email); err != nil {
		return err
	}

	return s.repository.CreateCustomer(customer)
}

func (s *OrderServiceImpl) CreateProduct(
	product model.Product,
) error {

	if err := utility.ValidateName(product.Name); err != nil {
		return err
	}

	if err := utility.ValidatePrice(product.Price); err != nil {
		return err
	}

	return s.repository.CreateProduct(product)
}

func (s *OrderServiceImpl) CreateOrder(
	customerID int,
) (int, error) {

	if customerID <= 0 {
		return 0, errors.New("customer ID must be greater than 0")
	}

	return s.repository.CreateOrder(customerID)
}

func (s *OrderServiceImpl) AddOrderItem(
	item model.OrderItem,
) error {

	if item.OrderID <= 0 {
		return errors.New("order ID must be greater than 0")
	}

	if item.ProductID <= 0 {
		return errors.New("product ID must be greater than 0")
	}

	if err := utility.ValidateQuantity(item.Quantity); err != nil {
		return err
	}

	if err := utility.ValidatePrice(item.Price); err != nil {
		return err
	}

	return s.repository.CreateOrderItem(item)
}

func (s *OrderServiceImpl) GetOrderDetails(
	orderID int,
) ([]model.OrderDetail, error) {

	if orderID <= 0 {
		return nil, errors.New("order ID must be greater than 0")
	}

	return s.repository.GetOrderDetails(orderID)
}
