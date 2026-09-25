package employee

import (
	"fmt"
	"companymanager/utils"
)

func PrintEmpInfo() {
	name := "Vimal Aditya"
	id := 101
	role := "Developer"
	salary := 50000

	utils.PrintHeader("Employee Details")

	fmt.Println("ID:", id)
	fmt.Println("Name:", name)
	fmt.Println("Role:", role)
	fmt.Println("Salary:", utils.Bonus(salary))
}