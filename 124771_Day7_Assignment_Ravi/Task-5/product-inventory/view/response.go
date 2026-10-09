package view

import "product-inventory/model"

type ProductResponse struct {
	Product model.Product
}

type ProductListResponse struct {
	Products []model.Product
}
