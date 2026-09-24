package employee

import "fmt"

func DeleteEmployee() {
	var id int

	fmt.Print("Enter employee ID to delete: ")
	fmt.Scan(&id)

	if _, exists := employeeMap[id]; !exists {
		fmt.Println("Employee not found.")
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}

	delete(employeeMap, id)

	fmt.Println("Employee deleted successfully.")
}
