package view

import (
	"employee-management/employeemanagement/controller"
	"fmt"
)

func ShowMessage() {
	message := controller.GetMessage()
	fmt.Println(message)
}