package service

import (
	"errors"
	"strings"

	"ecommerce-order-system/model"
	"ecommerce-order-system/repository"
)

type EcommerceServiceImpl struct {
	repository repository.EcommerceRepository
}

func NewEcommerceService(
	repository repository.EcommerceRepository,
) *EcommerceServiceImpl {
	return &EcommerceServiceImpl{
		repository: repository,
	}
}

func (s *EcommerceServiceImpl) CreateCustomer(
	customer model.Customer,
) error {
	if strings.TrimSpace(customer.Name) == "" {
		return errors.New("customer name cannot be empty")
	}

	if strings.TrimSpace(customer.Email) == "" {
		return errors.New("customer email cannot be empty")
	}

	return s.repository.CreateCustomer(customer)
}

func (s *EcommerceServiceImpl) CreateProduct(
	product model.Product,
) error {
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("product price cannot be negative")
	}

	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.repository.CreateProduct(product)
}

func (s *EcommerceServiceImpl) CreateOrder(
	order model.Order,
	items []model.OrderItem,
) (int, error) {
	if order.CustomerID <= 0 {
		return 0, errors.New("customer ID must be greater than zero")
	}

	if len(items) == 0 {
		return 0, errors.New("order must contain at least one item")
	}

	for _, item := range items {
		if item.ProductID <= 0 {
			return 0, errors.New("product ID must be greater than zero")
		}

		if item.Quantity <= 0 {
			return 0, errors.New("quantity must be greater than zero")
		}

		if item.Price < 0 {
			return 0, errors.New("item price cannot be negative")
		}
	}

	orderID, err := s.repository.CreateOrder(order)
	if err != nil {
		return 0, err
	}

	for _, item := range items {
		item.OrderID = orderID

		if err := s.repository.AddOrderItem(item); err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func (s *EcommerceServiceImpl) GetCustomers() ([]model.Customer, error) {
	return s.repository.GetCustomers()
}

func (s *EcommerceServiceImpl) GetProducts() ([]model.Product, error) {
	return s.repository.GetProducts()
}

func (s *EcommerceServiceImpl) GetOrderDetails(
	orderID int,
) ([]model.OrderDetails, error) {
	if orderID <= 0 {
		return nil, errors.New("order ID must be greater than zero")
	}

	return s.repository.GetOrderDetails(orderID)
}

func (s *EcommerceServiceImpl) GetAllOrderDetails() ([]model.OrderDetails, error) {
	return s.repository.GetAllOrderDetails()
}
