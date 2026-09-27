package main

import (
	"fmt"

	"employee-management-app/controller"
	"employee-management-app/repository"
	"employee-management-app/service"
	"employee-management-app/view"
)

func main() {

	// Create Repository.
	employeeRepository :=
		repository.NewEmployeeRepository()

	// Give Repository to Service.
	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	// Create View.
	employeeView :=
		view.NewEmployeeView()

	// Give Service and View to Controller.
	employeeController :=
		controller.NewEmployeeController(
			employeeService,
			employeeView,
		)

	fmt.Println(
		"Starting Employee Management System...",
	)

	// Start application.
	employeeController.Start()
}
