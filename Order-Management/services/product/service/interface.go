package service

import (
	"context"

	"order-management/services/product/model"
)

// ProductService defines the business operations for products.
type ProductService interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	Update(ctx context.Context, id uint, product *model.Product) (*model.Product, error)
	Delete(ctx context.Context, id uint) error
	ReserveStock(ctx context.Context, id uint, orderReference string, quantity int) (*model.Product, error)
	ReleaseStock(ctx context.Context, id uint, orderReference string, quantity int) error
}
