package main

import (
	"task01_employee_crud/config"
	"task01_employee_crud/controller"
	"task01_employee_crud/repository"
	"task01_employee_crud/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTable(db)

	repo := repository.NewEmployeeRepository(db)
	v := view.NewEmployeeView()
	ctrl := controller.NewEmployeeController(repo, v)

	ctrl.Run()
}
