package service

import (
	"errors"

	"ecommerce-management/model"
	"ecommerce-management/repository"
	"ecommerce-management/utility"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// Add Product
func (s *Service) AddProduct(product *model.Product) error {

	if err := utility.ValidateName(product.Name); err != nil {
		return err
	}

	if err := utility.ValidatePrice(product.Price); err != nil {
		return err
	}

	if err := utility.ValidateStock(product.Stock); err != nil {
		return err
	}

	if product.LowStockLimit < 0 {
		return errors.New("low stock limit cannot be negative")
	}

	return s.Repo.AddProduct(product)
}

// Get Products
func (s *Service) GetProducts() ([]model.Product, error) {

	return s.Repo.GetProducts()
}

// Get Product
func (s *Service) GetProduct(id int) (*model.Product, error) {

	if id <= 0 {
		return nil, errors.New("invalid product ID")
	}

	return s.Repo.GetProduct(id)
}

// Update Product
func (s *Service) UpdateProduct(product *model.Product) error {

	if product.ID <= 0 {
		return errors.New("invalid product ID")
	}

	if err := utility.ValidateName(product.Name); err != nil {
		return err
	}

	if err := utility.ValidatePrice(product.Price); err != nil {
		return err
	}

	if err := utility.ValidateStock(product.Stock); err != nil {
		return err
	}

	if product.LowStockLimit < 0 {
		return errors.New("low stock limit cannot be negative")
	}

	return s.Repo.UpdateProduct(product)
}

// Delete Product
func (s *Service) DeleteProduct(id int) error {

	if id <= 0 {
		return errors.New("invalid product ID")
	}

	return s.Repo.DeleteProduct(id)
}

// Increase Stock
func (s *Service) IncreaseStock(id int, quantity int) error {

	if id <= 0 {
		return errors.New("invalid product ID")
	}

	if err := utility.ValidateQuantity(quantity); err != nil {
		return err
	}

	return s.Repo.IncreaseStock(id, quantity)
}

// Decrease Stock
func (s *Service) DecreaseStock(id int, quantity int) error {

	if id <= 0 {
		return errors.New("invalid product ID")
	}

	if err := utility.ValidateQuantity(quantity); err != nil {
		return err
	}

	err := s.Repo.DecreaseStock(id, quantity)

	if err != nil {
		return errors.New("insufficient stock")
	}

	return nil
}

// Low Stock Products
func (s *Service) GetLowStockProducts() ([]model.Product, error) {

	return s.Repo.GetLowStockProducts()
}
