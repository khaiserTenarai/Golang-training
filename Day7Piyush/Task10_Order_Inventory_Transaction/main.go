package main

import (
	"task10_order_inventory_transaction/config"
	"task10_order_inventory_transaction/controller"
	"task10_order_inventory_transaction/repository"
	"task10_order_inventory_transaction/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	repo := repository.NewOrderRepository(db)
	v := view.NewOrderView()
	ctrl := controller.NewOrderController(repo, v)

	ctrl.Run()
}
