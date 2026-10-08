package main

import (
	"customer-management/controller"
	"customer-management/database"
	"customer-management/repository"
	"customer-management/service"
	"customer-management/view"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)
	customerView := view.NewCustomerView(customerController)

	customerView.Start()
}
