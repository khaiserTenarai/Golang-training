package employee

import "fmt"

func DisplayEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees available.")
		return
	}

	fmt.Println("<< Employee LIST >>")

	for _, emp := range employees {
		fmt.Println("ID:", emp.ID)
		fmt.Println("Name:", emp.Name)
		fmt.Println("Age:", emp.Age)
		fmt.Println("Salary:", emp.Salary)
	}
}
