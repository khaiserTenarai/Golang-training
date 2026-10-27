package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"order-management/services/product/model"
	"order-management/services/product/repository"

	"gorm.io/gorm"
)

// productService implements ProductService.
type productService struct {
	repo repository.ProductRepository
}

// NewProductService creates a ProductService.
func NewProductService(
	repo repository.ProductRepository,
) ProductService {
	return &productService{
		repo: repo,
	}
}

// Create validates and saves a new product.
func (s *productService) Create(
	ctx context.Context,
	product *model.Product,
) (*model.Product, error) {
	if product == nil {
		return nil, errors.New("product is required")
	}

	product.Name = strings.TrimSpace(product.Name)

	if err := validateProduct(product); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, fmt.Errorf("create product: %w", err)
	}

	return product, nil
}

// GetByID retrieves a product by ID.
func (s *productService) GetByID(
	ctx context.Context,
	id uint,
) (*model.Product, error) {
	if id == 0 {
		return nil, errors.New("product ID must be greater than zero")
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("get product: %w", err)
	}

	return product, nil
}

// GetAll retrieves all products.
func (s *productService) GetAll(
	ctx context.Context,
) ([]model.Product, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("get products: %w", err)
	}

	return products, nil
}

// Update validates and updates an existing product.
func (s *productService) Update(
	ctx context.Context,
	id uint,
	product *model.Product,
) (*model.Product, error) {
	if id == 0 {
		return nil, errors.New("product ID must be greater than zero")
	}

	if product == nil {
		return nil, errors.New("product is required")
	}

	product.Name = strings.TrimSpace(product.Name)

	if err := validateProduct(product); err != nil {
		return nil, err
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("find product for update: %w", err)
	}

	existing.Name = product.Name
	existing.Description = product.Description
	existing.Price = product.Price
	existing.Stock = product.Stock

	if err := s.repo.Update(ctx, existing); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("update product: %w", err)
	}

	return existing, nil
}

// Delete removes an existing product.
func (s *productService) Delete(
	ctx context.Context,
	id uint,
) error {
	if id == 0 {
		return errors.New("product ID must be greater than zero")
	}

	if _, err := s.repo.GetByID(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return fmt.Errorf("find product for deletion: %w", err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete product: %w", err)
	}

	return nil
}

// validateProduct checks the business rules for product data.
func validateProduct(product *model.Product) error {
	if product == nil {
		return errors.New("product is required")
	}

	if product.Name == "" {
		return errors.New("product name is required")
	}

	if len(product.Name) > 150 {
		return errors.New("product name must not exceed 150 characters")
	}

	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}

	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return nil
}

func (s *productService) ReserveStock(ctx context.Context, id uint, orderReference string, quantity int) (*model.Product, error) {
	if id == 0 || strings.TrimSpace(orderReference) == "" || quantity <= 0 {
		return nil, errors.New("product ID, order reference and positive quantity are required")
	}
	product, err := s.repo.ReserveStock(ctx, id, orderReference, quantity)
	if err != nil {
		return nil, fmt.Errorf("reserve product stock: %w", err)
	}
	return product, nil
}

func (s *productService) ReleaseStock(ctx context.Context, id uint, orderReference string, quantity int) error {
	if id == 0 || strings.TrimSpace(orderReference) == "" || quantity <= 0 {
		return errors.New("product ID, order reference and positive quantity are required")
	}
	if err := s.repo.ReleaseStock(ctx, id, orderReference, quantity); err != nil {
		return fmt.Errorf("release product stock: %w", err)
	}
	return nil
}
