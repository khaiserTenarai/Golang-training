package employee

import "fmt"

func SearchEmployee() {
	var id int

	fmt.Print("Enter employee ID to search: ")
	fmt.Scan(&id)

	emp, exists := employeeMap[id]

	if !exists {
		fmt.Println("Employee not found.")
		return
	}

	fmt.Println("Employee Details")
	fmt.Println("ID:", emp.ID)
	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Salary:", emp.Salary)
}
