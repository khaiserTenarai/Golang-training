package main

import (
	"employee-management/controller"
	"employee-management/database"
	"employee-management/repository"
	"employee-management/service"
	"employee-management/view"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()
	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)
	employeeView := view.NewEmployeeView(employeeController)

	salaryRepository := repository.NewSalaryRepository(db)
	salaryService := service.NewSalaryService(salaryRepository)
	salaryController := controller.NewSalaryController(salaryService)
	salaryView := view.NewSalaryView(salaryController)

	employeeView.Start()
	//task4
	salaryView.UpdateSalary()
}
