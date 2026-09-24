// TASK 6
/*package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Employee Management System")

	id := uuid.New()

	fmt.Println("Employee ID:", id)
} */



// TASK 7
/*package main

import (
	"fmt"
	"os"
	"employeeapp/internal/employee"
	"employeeapp/pkg/utils"
)

func main() {

	emp := employee.NewEmployee(101, "Piyush", 50000)

	fmt.Println("Employee ID:", emp.ID)
	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("Monthly Salary:", emp.Salary)

	annualSalary := utils.CalculateAnnualSalary(emp.Salary)

	fmt.Println("Annual Salary:", annualSalary)

	// Read environment variable(TASK 10)
	appName := os.Getenv("APP_NAME")

	if appName == "" {
		fmt.Println("APP_NAME is not set")
	} else {
		fmt.Println("Application Name:", appName)
	}

} */


// Task 11
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"employeeapp/internal/employee"
// )

// func main() {
// 	// Environment variable
// 	appName := os.Getenv("APP_NAME")

// 	if appName == "" {
// 		fmt.Println("APP_NAME is not set")
// 	} else {
// 		fmt.Println("Application Name:", appName)
// 	}

// 	// Employee data
// 	employees := []employee.Employee{
// 		employee.NewEmployee(101, "Piyush", 50000),
// 		employee.NewEmployee(102, "Rahul", 60000),
// 		employee.NewEmployee(103, "Amit", 55000),
// 	}

// 	// Get employee ID from command line
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide employee ID")
// 		fmt.Println("Example: go run . 102")
// 		return
// 	}

// 	id, err := strconv.Atoi(os.Args[1])

// 	if err != nil {
// 		fmt.Println("Invalid employee ID")
// 		return
// 	}

// 	// Search employee
// 	for _, emp := range employees {
// 		if emp.ID == id {
// 			fmt.Println("Employee Found!")
// 			fmt.Println("ID:", emp.ID)
// 			fmt.Println("Name:", emp.Name)
// 			fmt.Println("Salary:", emp.Salary)
// 			return
// 		}
// 	}

// 	fmt.Println("Employee not found")
// }

// TASK 12 
/*
cmd → Entry point of the code(main.go)
internal → code which is used inside the application
pkg → reusable packages/utilities
*/

// Task 15

package main

import (
	"fmt"

	"employeeapp/internal/employee"
	"employeeapp/pkg/utils"
)

func main() {

	employees := []employee.Employee{
		employee.NewEmployee(101, "Piyush", 50000),
		employee.NewEmployee(102, "Rahul", 60000),
		employee.NewEmployee(103, "Amit", 55000),
	}

	for {
		fmt.Println()
		fmt.Println("===== Employee Management System =====")
		fmt.Println("1. Show All Employees")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Calculate Annual Salary")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {

			fmt.Println("\nEmployee List:")

			for _, emp := range employees {
				fmt.Println("ID:", emp.ID)
				fmt.Println("Name:", emp.Name)
				fmt.Println("Salary:", emp.Salary)
				fmt.Println()
			}

		} else if choice == 2 {

			var id int

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for _, emp := range employees {
				if emp.ID == id {
					fmt.Println("\nEmployee Found!")
					fmt.Println("ID:", emp.ID)
					fmt.Println("Name:", emp.Name)
					fmt.Println("Salary:", emp.Salary)

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found")
			}

		} else if choice == 3 {

			var id int

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			found := false

			for _, emp := range employees {
				if emp.ID == id {

					annualSalary := utils.CalculateAnnualSalary(emp.Salary)

					fmt.Println("Employee:", emp.Name)
					fmt.Println("Annual Salary:", annualSalary)

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found")
			}

		} else if choice == 4 {

			fmt.Println("Exiting Employee Management System...")
			break

		} else {

			fmt.Println("Invalid choice")
		}
	}
}