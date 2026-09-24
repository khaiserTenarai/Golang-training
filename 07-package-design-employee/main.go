package main

import (
	"07-package-design-employee/employee"
	"07-package-design-employee/utils"
)

func main(){
	name :=  employee.GetEmployeeName()
	utils.PrintEmployeeName(name)
}

/*
output :
------------
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment\07-package-design-employee> go run .
 employee name :  Ganesh
 */
 