package main

import (
	"task02_department_management/config"
	"task02_department_management/controller"
	"task02_department_management/repository"
	"task02_department_management/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	deptRepo := repository.NewDepartmentRepository(db)
	empRepo := repository.NewEmployeeRepository(db)
	v := view.NewDepartmentView()
	ctrl := controller.NewDepartmentController(deptRepo, empRepo, v)

	ctrl.Run()
}
