package dto

import (
	"errors"
	"strings"
)

// CreateProductRequest contains the data needed to create a product.
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}

// UpdateProductRequest contains the data used to update a product.
type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}

// Validate checks product data before it reaches the service layer.
func (r CreateProductRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("product name is required")
	}

	if len(r.Name) > 150 {
		return errors.New("product name must not exceed 150 characters")
	}

	if r.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}

	if r.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return nil
}

// Validate checks product update data.
func (r UpdateProductRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("product name is required")
	}

	if len(r.Name) > 150 {
		return errors.New("product name must not exceed 150 characters")
	}

	if r.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}

	if r.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return nil
}
