package view

import "product-inventory/model"

type ProductView interface {
	Start()
	CreateProduct()
	GetProduct()
	GetAllProducts()
	UpdateProduct()
	DeleteProduct()
	IncreaseStock()
	DecreaseStock()
	GetLowStockProducts()
	DisplayProduct(product *model.Product)
	DisplayProducts(products []model.Product)
}
