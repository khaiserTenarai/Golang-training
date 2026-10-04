package main

import (
	"task09_ecommerce_order_system/config"
	"task09_ecommerce_order_system/controller"
	"task09_ecommerce_order_system/repository"
	"task09_ecommerce_order_system/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	custRepo := repository.NewCustomerRepository(db)
	prodRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	v := view.NewOrderView()
	ctrl := controller.NewOrderController(custRepo, prodRepo, orderRepo, v)

	ctrl.Run()
}
