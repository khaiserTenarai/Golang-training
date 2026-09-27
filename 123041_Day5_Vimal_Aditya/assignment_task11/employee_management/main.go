package main

import (
	"fmt"
	"employee_management/dao"
	"employee_management/service"
	"employee_management/controller"
	"employee_management/view"
)

func main(){
	fmt.Println("Employee Management System: ")

	repo := dao.NewEmployeeDao()
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)
	ui := view.NewEmployeeView(ctrl)

	ui.Start()

}

