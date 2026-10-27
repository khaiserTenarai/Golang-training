package repository

import (
	"context"
	"order-management/services/order/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	GetByID(ctx context.Context, id uint) (*model.Order, error)
	GetAll(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	SetStatus(ctx context.Context, id uint, newStatus string, allowedFrom []string) (bool, error)
	MarkStockReleased(ctx context.Context, id uint) error
}
