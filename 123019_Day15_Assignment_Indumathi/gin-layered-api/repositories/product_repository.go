package repositories

import "gin-layered-api/models"

type ProductRepository interface {
	GetAll(page, limit int) ([]models.Product, int)
	Create(product *models.Product) *models.Product
}

type productRepository struct {
	products []models.Product
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		products: []models.Product{
			{ID: 1, Name: "Laptop", Price: 999.99},
			{ID: 2, Name: "Mouse", Price: 25.50},
			{ID: 3, Name: "Keyboard", Price: 75.00},
			{ID: 4, Name: "Monitor", Price: 200.00},
			{ID: 5, Name: "Headphones", Price: 50.00},
		},
	}
}

func (r *productRepository) GetAll(page, limit int) ([]models.Product, int) {
	total := len(r.products)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= total {
		return []models.Product{}, total
	}

	if endIndex > total {
		endIndex = total
	}

	return r.products[startIndex:endIndex], total
}

func (r *productRepository) Create(product *models.Product) *models.Product {
	product.ID = len(r.products) + 1
	r.products = append(r.products, *product)
	return product
}
