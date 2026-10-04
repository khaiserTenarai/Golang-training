package main

import (
	"emp_management_Updated/controller"
	"emp_management_Updated/dao"
	"emp_management_Updated/service"
	"emp_management_Updated/view"
)

func main() {

	// Create DAO
	employeeDAO := dao.NewEmployeeDAO()

	// Create Service
	employeeService := service.NewEmployeeService(employeeDAO)

	// Create Controller
	employeeController := controller.NewEmployeeController(employeeService)

	// Create View
	employeeView := view.NewEmployeeView(employeeController)

	// Start application
	employeeView.ShowMenu()
}
