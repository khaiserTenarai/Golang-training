package main

import (
	"employee-management/controller"
	"employee-management/database"
	"employee-management/repository"
	"employee-management/service"
	"employee-management/view"
)

// Task2
func main() {
	db := database.ConnectDB()
	defer db.Close()
	departmentRepository := repository.NewDepartmentRepository(db)
	departmentService := service.NewDepartmentService(departmentRepository)
	departmentController := controller.NewDepartmentController(departmentService)
	departmentView := view.NewDepartmentView(departmentController)
	departmentView.Start()
}
