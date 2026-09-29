package main

import (
	"employee-management/controller"
	"employee-management/dao"
	"employee-management/service"
	"employee-management/view"
)

func main() {

	// Create DAO
	employeeDAO := dao.NewEmployeeDAO()

	// Create Service
	employeeService := service.NewEmployeeService(
		employeeDAO,
	)

	// Create Controller
	employeeController := controller.NewEmployeeController(
		employeeService,
	)

	// Create View
	employeeView := view.NewEmployeeView(
		employeeController,
	)

	// Start application
	employeeView.Start()
}
