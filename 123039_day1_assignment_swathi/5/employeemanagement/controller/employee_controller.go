package controller

import "employee-management/employeemanagement/service"

func GetMessage() string{
	return service.GetMessage()
}