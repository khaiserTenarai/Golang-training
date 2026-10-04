package service

import (
	"errors"
	"strings"

	"prodcut-inventory/model"
	"prodcut-inventory/repository"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(
	repository repository.ProductRepository,
) ProductService {

	return &ProductServiceImpl{
		repository: repository,
	}
}

// CREATE
func (s *ProductServiceImpl) AddProduct(
	product model.Product,
) error {

	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	if product.LowStockLimit < 0 {
		return errors.New("low stock limit cannot be negative")
	}

	return s.repository.Save(product)
}

// READ ONE
func (s *ProductServiceImpl) FindProductByID(
	id int,
) (model.Product, error) {

	if id <= 0 {
		return model.Product{}, errors.New(
			"ID must be greater than 0",
		)
	}

	return s.repository.FindByID(id)
}

// READ ALL
func (s *ProductServiceImpl) FindAllProducts() []model.Product {

	return s.repository.FindAll()
}

// UPDATE
func (s *ProductServiceImpl) UpdateProduct(
	product model.Product,
) error {

	if product.ID <= 0 {
		return errors.New("ID must be greater than 0")
	}

	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	if product.LowStockLimit < 0 {
		return errors.New("low stock limit cannot be negative")
	}

	return s.repository.Update(product)
}

// DELETE
func (s *ProductServiceImpl) DeleteProduct(
	id int,
) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	return s.repository.Delete(id)
}

// INCREASE STOCK
func (s *ProductServiceImpl) IncreaseStock(
	id int,
	quantity int,
) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return s.repository.IncreaseStock(
		id,
		quantity,
	)
}

// DECREASE STOCK
func (s *ProductServiceImpl) DecreaseStock(
	id int,
	quantity int,
) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return s.repository.DecreaseStock(
		id,
		quantity,
	)
}

// LOW STOCK
func (s *ProductServiceImpl) FindLowStock() []model.Product {

	return s.repository.FindLowStock()
}
