package service

import (
	"order-inventory/repository"
	"order-inventory/utility"
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

func (s *OrderServiceImpl) CreateOrder(
	customerName string,
	productID int,
	quantity int,
) error {

	// Validation

	if err := utility.ValidateCustomerName(
		customerName,
	); err != nil {
		return err
	}

	if err := utility.ValidateProductID(
		productID,
	); err != nil {
		return err
	}

	if err := utility.ValidateQuantity(
		quantity,
	); err != nil {
		return err
	}

	// Repository

	return s.repository.CreateOrder(
		customerName,
		productID,
		quantity,
	)
}
