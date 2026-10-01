package service

import "example.com/employee-management/model"

type ProductService interface {
    AddProduct(p model.Product) error
    GetProduct(id int64) (model.Product, error)
    GetAllProducts() ([]model.Product, error)
    UpdateProduct(p model.Product) error
    DeleteProduct(id int64) error
    IncreaseStock(id int64, amount int) error
    DecreaseStock(id int64, amount int) error
    GetLowStockProducts() ([]model.Product, error)
}