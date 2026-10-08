package service

import (
	"product-inventory/model"
	"product-inventory/repository"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		repository: repository,
	}
}

func (s *ProductServiceImpl) CreateProduct(product model.Product) error {
	return s.repository.CreateProduct(product)
}

func (s *ProductServiceImpl) GetProduct(id int) (*model.Product, error) {
	return s.repository.GetProduct(id)
}

func (s *ProductServiceImpl) GetAllProducts() ([]model.Product, error) {
	return s.repository.GetAllProducts()
}

func (s *ProductServiceImpl) UpdateProduct(product model.Product) error {
	return s.repository.UpdateProduct(product)
}

func (s *ProductServiceImpl) DeleteProduct(id int) error {
	return s.repository.DeleteProduct(id)
}

func (s *ProductServiceImpl) IncreaseStock(id int, quantity int) error {
	return s.repository.IncreaseStock(id, quantity)
}

func (s *ProductServiceImpl) DecreaseStock(id int, quantity int) error {
	return s.repository.DecreaseStock(id, quantity)
}

func (s *ProductServiceImpl) GetLowStockProducts() ([]model.Product, error) {
	return s.repository.GetLowStockProducts()
}
