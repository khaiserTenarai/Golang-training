package employee

import "fmt"

func AddEmployee() {
	var emp Employee

	fmt.Print("Enter employee ID: ")
	fmt.Scan(&emp.ID)

	if _, exists := employeeMap[emp.ID]; exists {
		fmt.Println("Employee ID already exists.")
		return
	}

	fmt.Print("Enter employee name: ")
	fmt.Scan(&emp.Name)

	fmt.Print("Enter employee age: ")
	fmt.Scan(&emp.Age)

	fmt.Print("Enter employee salary: ")
	fmt.Scan(&emp.Salary)

	employees = append(employees, emp)
	employeeMap[emp.ID] = emp

	fmt.Println("Employee added successfully.")

}
