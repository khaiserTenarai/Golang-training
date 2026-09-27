package main

import (
	"emp_simple/controller"
	"emp_simple/service"
	"emp_simple/view"
)

func main() {
	empService := &service.EmpService{}

	controller := &controller.EmployeeController{
		Service: empService,
	}

	view := &view.EmpView{
		Controller: controller,
	}

	view.ShowMenu()
}
