package controller

import (
	"product-inventory/model"
	"product-inventory/service"
)

type ProductControllerImpl struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &ProductControllerImpl{
		service: service,
	}
}

func (c *ProductControllerImpl) CreateProduct(product model.Product) error {
	return c.service.CreateProduct(product)
}

func (c *ProductControllerImpl) GetProduct(id int) (*model.Product, error) {
	return c.service.GetProduct(id)
}

func (c *ProductControllerImpl) GetAllProducts() ([]model.Product, error) {
	return c.service.GetAllProducts()
}

func (c *ProductControllerImpl) UpdateProduct(product model.Product) error {
	return c.service.UpdateProduct(product)
}

func (c *ProductControllerImpl) DeleteProduct(id int) error {
	return c.service.DeleteProduct(id)
}

func (c *ProductControllerImpl) IncreaseStock(id int, quantity int) error {
	return c.service.IncreaseStock(id, quantity)
}

func (c *ProductControllerImpl) DecreaseStock(id int, quantity int) error {
	return c.service.DecreaseStock(id, quantity)
}

func (c *ProductControllerImpl) GetLowStockProducts() ([]model.Product, error) {
	return c.service.GetLowStockProducts()
}
