package services

import (
	"gin-layered-api/models"
	"gin-layered-api/repositories"
)

type ProductService interface {
	GetProducts(page, limit int) models.PaginatedProductsResponse
	CreateProduct(req models.CreateProductRequest) *models.Product
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) GetProducts(page, limit int) models.PaginatedProductsResponse {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 2
	}

	data, total := s.productRepo.GetAll(page, limit)

	return models.PaginatedProductsResponse{
		Page:  page,
		Limit: limit,
		Total: total,
		Data:  data,
	}
}

func (s *productService) CreateProduct(req models.CreateProductRequest) *models.Product {
	product := &models.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	return s.productRepo.Create(product)
}
