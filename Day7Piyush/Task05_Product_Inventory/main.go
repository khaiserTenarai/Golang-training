package main

import (
	"task05_product_inventory/config"
	"task05_product_inventory/controller"
	"task05_product_inventory/repository"
	"task05_product_inventory/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTable(db)

	repo := repository.NewProductRepository(db)
	v := view.NewProductView()
	ctrl := controller.NewProductController(repo, v)

	ctrl.Run()
}
