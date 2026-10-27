package dto

import (
	"errors"
)

type CreateOrderRequest struct {
	CustomerID uint `json:"customerId" binding:"required"`
	ProductID  uint `json:"productId" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}

func (r CreateOrderRequest) Validate() error {
	if r.CustomerID == 0 {
		return errors.New("customerId must be greater than zero")
	}
	if r.ProductID == 0 {
		return errors.New("productId must be greater than zero")
	}
	if r.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	return nil
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
