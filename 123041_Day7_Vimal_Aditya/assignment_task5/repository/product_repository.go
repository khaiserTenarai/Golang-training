package repository

import (
    "errors"

    "example.com/employee-management/model"
)

var (
    ErrProductNotFound   = errors.New("product not found")
    ErrInsufficientStock = errors.New("insufficient stock available")
)

type ProductRepository interface {
    Save(product model.Product) error
    FindByID(id int64) (model.Product, error)
    FindAll() ([]model.Product, error)
    Update(product model.Product) error
    Delete(id int64) error
    AdjustStock(id int64, amount int) error
    FindLowStockProducts() ([]model.Product, error)
}