package main

import (
	"task04_salary_management/config"
	"task04_salary_management/controller"
	"task04_salary_management/repository"
	"task04_salary_management/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	repo := repository.NewSalaryRepository(db)
	v := view.NewSalaryView()
	ctrl := controller.NewSalaryController(repo, v)

	ctrl.Run()
}
