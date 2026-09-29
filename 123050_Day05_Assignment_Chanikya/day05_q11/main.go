package main

import (
	"employee-management/controller"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
)

func main() {

	employee := model.Employee{
		ID:     101,
		Name:   "Ram",
		Salary: 50000,
	}

	repo := repository.NewEmployeeRepository()

	employeeService := service.NewEmployeeService(repo)

	employeeController := controller.NewEmployeeController(employeeService)

	employeeController.AddEmployee(employee)
	employeeController.GetEmployee(101)
}
