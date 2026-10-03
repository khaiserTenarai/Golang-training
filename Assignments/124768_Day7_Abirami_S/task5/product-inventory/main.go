package main

import (
	"product-inventory/controller"
	"product-inventory/database"
	"product-inventory/repository"
	"product-inventory/service"
	"product-inventory/view"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	productView := view.NewProductView(productController)

	productView.Start()
}
