package main

import (
	"task06_customer_management/config"
	"task06_customer_management/controller"
	"task06_customer_management/repository"
	"task06_customer_management/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTable(db)

	repo := repository.NewCustomerRepository(db)
	v := view.NewCustomerView()
	ctrl := controller.NewCustomerController(repo, v)

	ctrl.Run()
}
