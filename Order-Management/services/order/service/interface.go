package service

import (
	"context"

	"order-management/services/order/model"
)

type OrderService interface {
	Create(
		ctx context.Context,
		customerID uint,
		productID uint,
		quantity int,
	) (*model.Order, error)

	GetByID(ctx context.Context, id uint) (*model.Order, error)
	GetAll(ctx context.Context) ([]model.Order, error)

	UpdateStatus(
		ctx context.Context,
		id uint,
		status string,
	) (*model.Order, error)
}
