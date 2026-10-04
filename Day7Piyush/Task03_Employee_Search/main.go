package main

import (
	"task03_employee_search/config"
	"task03_employee_search/controller"
	"task03_employee_search/repository"
	"task03_employee_search/view"
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
