package service

import (
    "errors"
    "strings"

    "example.com/employee-management/model"
    "example.com/employee-management/repository"
)

type productServiceImpl struct {
    repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
    return &productServiceImpl{repo: repo}
}

func (s *productServiceImpl) validate(p model.Product) error {
    if strings.TrimSpace(p.Name) == "" {
        return errors.New("product name is required")
    }
    if strings.TrimSpace(p.Category) == "" {
        return errors.New("category is required")
    }
    if p.Price < 0 {
        return errors.New("price cannot be negative")
    }
    if p.StockQuantity < 0 {
        return errors.New("stock quantity cannot be negative")
    }
    if p.ReorderLevel < 0 {
        return errors.New("reorder level cannot be negative")
    }
    return nil
}

func (s *productServiceImpl) AddProduct(p model.Product) error {
    if err := s.validate(p); err != nil {
        return err
    }
    return s.repo.Save(p)
}

func (s *productServiceImpl) GetProduct(id int64) (model.Product, error) {
    if id <= 0 {
        return model.Product{}, errors.New("invalid product ID")
    }
    return s.repo.FindByID(id)
}

func (s *productServiceImpl) GetAllProducts() ([]model.Product, error) {
    return s.repo.FindAll()
}

func (s *productServiceImpl) UpdateProduct(p model.Product) error {
    if p.ID <= 0 {
        return errors.New("invalid product ID")
    }
    if err := s.validate(p); err != nil {
        return err
    }
    return s.repo.Update(p)
}

func (s *productServiceImpl) DeleteProduct(id int64) error {
    if id <= 0 {
        return errors.New("invalid product ID")
    }
    return s.repo.Delete(id)
}

func (s *productServiceImpl) IncreaseStock(id int64, amount int) error {
    if id <= 0 {
        return errors.New("invalid product ID")
    }
    if amount <= 0 {
        return errors.New("increase amount must be greater than zero")
    }
    return s.repo.AdjustStock(id, amount)
}

func (s *productServiceImpl) DecreaseStock(id int64, amount int) error {
    if id <= 0 {
        return errors.New("invalid product ID")
    }
    if amount <= 0 {
        return errors.New("decrease amount must be greater than zero")
    }
    return s.repo.AdjustStock(id, -amount)
}

func (s *productServiceImpl) GetLowStockProducts() ([]model.Product, error) {
    return s.repo.FindLowStockProducts()
}