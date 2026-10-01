// This program walks through Department Management, Employee Search,
// and Salary Management end to end, against a real Postgres database.
//
// Before running this, make sure:
//  1. Postgres is running and a database exists (see README.md).
//  2. sql/schema.sql has been applied to that database.
//  3. DB_HOST / DB_PORT / DB_USER / DB_PASSWORD / DB_NAME are set as
//     needed (or the defaults in db/db.go match your setup).
package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"employee-department-service/db"
	"employee-department-service/model"
	"employee-department-service/repository"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("could not connect to the database: ", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		choice := readInt(reader, "Choose an operation: ")

		var operationErr error
		switch choice {
		case 1:
			name := readString(reader, "Department name: ")
			var id int
			id, operationErr = repository.CreateDepartment(conn, name)
			if operationErr == nil {
				fmt.Println("Created department with ID:", id)
			}
		case 2:
			var departments []model.Department
			departments, operationErr = repository.ListDepartments(conn)
			for _, department := range departments {
				fmt.Printf("[%d] %s\n", department.ID, department.Name)
			}
		case 3:
			id := readInt(reader, "Department ID: ")
			var department model.Department
			department, operationErr = repository.GetDepartment(conn, id)
			if operationErr == nil {
				fmt.Printf("[%d] %s\n", department.ID, department.Name)
			}
		case 4:
			id := readInt(reader, "Department ID: ")
			name := readString(reader, "New department name: ")
			operationErr = repository.UpdateDepartment(conn, id, name)
		case 5:
			id := readInt(reader, "Department ID: ")
			operationErr = repository.DeleteDepartment(conn, id)
		case 6:
			departmentID := readOptionalInt(reader, "Department ID (press Enter for none): ")
			employee := model.Employee{
				Name:         readString(reader, "Employee name: "),
				Email:        readString(reader, "Employee email: "),
				Age:          readInt(reader, "Employee age: "),
				Salary:       readFloat(reader, "Employee salary: "),
				DepartmentID: departmentID,
			}
			var id int
			id, operationErr = repository.CreateEmployee(conn, employee)
			if operationErr == nil {
				fmt.Println("Created employee with ID:", id)
			}
		case 7:
			id := readInt(reader, "Employee ID: ")
			var employee model.Employee
			employee, operationErr = repository.GetEmployee(conn, id)
			if operationErr == nil {
				fmt.Printf("[%d] %s, %s, age %d, salary %.2f\n", employee.ID, employee.Name, employee.Email, employee.Age, employee.Salary)
			}
		case 8:
			id := readInt(reader, "Employee ID: ")
			operationErr = repository.DeleteEmployee(conn, id)
		case 9:
			results, total, searchErr := searchEmployees(reader, conn)
			operationErr = searchErr
			if operationErr == nil {
				fmt.Println("Total matches:", total)
				for _, employee := range results {
					fmt.Printf("[%d] %s - salary %.2f\n", employee.ID, employee.Name, employee.Salary)
				}
			}
		case 10:
			id := readInt(reader, "Employee ID: ")
			newSalary := readFloat(reader, "New salary: ")
			operationErr = repository.UpdateEmployeeSalary(conn, id, newSalary)
		case 11:
			id := readInt(reader, "Employee ID: ")
			var history []model.SalaryHistory
			history, operationErr = repository.GetSalaryHistory(conn, id)
			for _, change := range history {
				fmt.Printf("%.2f -> %.2f at %s\n", change.OldSalary, change.NewSalary, change.ChangedAt.Format("2006-01-02 15:04:05"))
			}
		case 0:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("Unknown operation.")
			continue
		}

		if operationErr != nil {
			if errors.Is(operationErr, repository.ErrDepartmentHasEmployees) {
				fmt.Println("Cannot delete department:", operationErr)
			} else {
				fmt.Println("Operation failed:", operationErr)
			}
		} else if choice >= 1 && choice <= 11 && choice != 2 && choice != 3 && choice != 6 && choice != 7 && choice != 9 && choice != 11 {
			fmt.Println("Operation completed successfully.")
		}
	}
}

func printMenu() {
	fmt.Println("\n--- Employee Department Service ---")
	fmt.Println("1. Create department")
	fmt.Println("2. List departments")
	fmt.Println("3. Get department")
	fmt.Println("4. Update department")
	fmt.Println("5. Delete department")
	fmt.Println("6. Create employee")
	fmt.Println("7. Get employee")
	fmt.Println("8. Delete employee")
	fmt.Println("9. Search employees")
	fmt.Println("10. Update employee salary")
	fmt.Println("11. Get salary history")
	fmt.Println("0. Exit")
}

func searchEmployees(reader *bufio.Reader, conn *sql.DB) ([]model.Employee, int, error) {
	params := repository.EmployeeSearchParams{
		Name:           readString(reader, "Name filter (press Enter for any): "),
		DepartmentName: readString(reader, "Department filter (press Enter for any): "),
		MinSalary:      readOptionalFloat(reader, "Minimum salary (press Enter for none): "),
		MaxSalary:      readOptionalFloat(reader, "Maximum salary (press Enter for none): "),
		SortBy:         readString(reader, "Sort by name, salary, age, or id: "),
		SortOrder:      readString(reader, "Sort order asc or desc: "),
		Page:           readInt(reader, "Page number: "),
		PageSize:       readInt(reader, "Page size: "),
	}
	return repository.SearchEmployees(conn, params)
}

func readString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		value := readString(reader, prompt)
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
		fmt.Println("Please enter a whole number.")
	}
}

func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		value := readString(reader, prompt)
		parsed, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return parsed
		}
		fmt.Println("Please enter a number.")
	}
}

func readOptionalInt(reader *bufio.Reader, prompt string) *int {
	for {
		value := readString(reader, prompt)
		if value == "" {
			return nil
		}
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return &parsed
		}
		fmt.Println("Please enter a whole number or press Enter.")
	}
}

func readOptionalFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		value := readString(reader, prompt)
		if value == "" {
			return 0
		}
		parsed, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return parsed
		}
		fmt.Println("Please enter a number or press Enter.")
	}
}
