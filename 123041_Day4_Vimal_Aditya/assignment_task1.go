package main

import (
	"errors"
	"fmt"
)

var ids []int
var names []string
var salaries []float64

func createEmployee(id int, name string, salary float64) error {
	for _, existingID := range ids {
		if existingID == id {
			return errors.New("employee ID already exists")
		}
	}

	ids = append(ids, id)
	names = append(names, name)
	salaries = append(salaries, salary)
	return nil
}

func getEmployee(id int) (*string, *float64, error) {
	for i, existingID := range ids {
		if existingID == id {
			return &names[i], &salaries[i], nil
		}
	}
	return nil, nil, errors.New("employee not found")
}

func updateEmployee(id int, newName string, newSalary float64) error {
	namePtr, salaryPtr, err := getEmployee(id)
	if err != nil {
		return err
	}

	*namePtr = newName
	*salaryPtr = newSalary
	return nil
}

// DELETE
func deleteEmployee(id int) error {
	for i, existingID := range ids {
		if existingID == id {
			ids = append(ids[:i], ids[i+1:]...)
			names = append(names[:i], names[i+1:]...)
			salaries = append(salaries[:i], salaries[i+1:]...)
			return nil
		}
	}
	return errors.New("employee not found")
}

func main() {

	fmt.Println("\n************************************")
	fmt.Println("1. Create functions for employee CRUD.")
	fmt.Println("**************************************")


	var choice int

	for {
		fmt.Println("\n--- EMPLOYEE MANAGEMENT MENU ---")
		fmt.Println("1. Create Employee")
		fmt.Println("2. View Employee")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. View All Employees")
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var name string
			var salary float64

			fmt.Print("Enter ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Salary: ")
			fmt.Scan(&salary)

			err := createEmployee(id, name, salary)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee created successfully!")
			}

		case 2:
			var id int
			fmt.Print("Enter ID to search: ")
			fmt.Scan(&id)

			namePtr, salaryPtr, err := getEmployee(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Found -> Name:", *namePtr, "| Salary:", *salaryPtr)
			}

		case 3:
			var id int
			var newName string
			var newSalary float64

			fmt.Print("Enter ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Name: ")
			fmt.Scan(&newName)
			fmt.Print("Enter New Salary: ")
			fmt.Scan(&newSalary)

			err := updateEmployee(id, newName, newSalary)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee updated successfully!")
			}

		case 4:
			var id int
			fmt.Print("Enter ID to delete: ")
			fmt.Scan(&id)

			err := deleteEmployee(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee deleted successfully!")
			}

		case 5:
			if len(ids) == 0 {
				fmt.Println("No employee records found.")
			} else {
				fmt.Println("\n--- All Employees ---")
				for i := 0; i < len(ids); i++ {
					fmt.Println("ID:", ids[i], "| Name:", names[i], "| Salary:", salaries[i])
				}
			}

		case 6:
			fmt.Println("Thank u")
			return

		default:
			fmt.Println("invalid choice, try again")
		}
	}
}