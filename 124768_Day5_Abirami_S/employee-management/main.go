package main

import (
	"employee-management/controller"
	"employee-management/repository"
	"employee-management/service"
	"employee-management/view"
)

func main() {
	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)
	employeeView := view.NewEmployeeView(employeeController)

	employeeView.Start()
}
