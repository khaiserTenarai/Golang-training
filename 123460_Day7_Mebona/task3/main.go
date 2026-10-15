package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	connString := "postgres://postgres:pgadmin@localhost:5432/employee_db"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	for {
		fmt.Println("\n==============================")
		fmt.Println("       EMPLOYEE SEARCH")
		fmt.Println("==============================")
		fmt.Println("1. Search Employees (by Name, Department, Salary)")
		fmt.Println("2. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}

		switch choice {
		case 1:
			searchEmployees(conn)
		case 2:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select 1 or 2.")
		}
	}
}

func searchEmployees(conn *pgx.Conn) {
	fmt.Println("\n----- SEARCH EMPLOYEES -----")
	
	nameQuery := readInput("Filter by Name (leave blank for any): ")
	deptQuery := readInput("Filter by Department Name (leave blank for any): ")
	
	minSalaryStr := readInput("Minimum Salary (leave blank for none): ")
	var minSalary float64 = 0
	if minSalaryStr != "" {
		if val, err := strconv.ParseFloat(minSalaryStr, 64); err == nil {
			minSalary = val
		}
	}

	maxSalaryStr := readInput("Maximum Salary (leave blank for none): ")
	var maxSalary float64 = 999999999
	if maxSalaryStr != "" {
		if val, err := strconv.ParseFloat(maxSalaryStr, 64); err == nil {
			maxSalary = val
		}
	}

	fmt.Println("\nSort options: 1. ID  2. Name  3. Salary  4. Age")
	sortChoice := readInput("Choose sort column [1-4, default 1]: ")
	
	sortColumn := "e.id"
	switch sortChoice {
	case "2":
		sortColumn = "e.name"
	case "3":
		sortColumn = "e.salary"
	case "4":
		sortColumn = "e.age"
	}

	sortDir := readInput("Sort direction (ASC / DESC, default ASC): ")
	if strings.ToUpper(sortDir) != "DESC" {
		sortDir = "ASC"
	} else {
		sortDir = "DESC"
	}

	limitStr := readInput("Items per page (pagination limit, default 5): ")
	limit := 5
	if val, err := strconv.Atoi(limitStr); err == nil && val > 0 {
		limit = val
	}

	pageStr := readInput("Page number (default 1): ")
	page := 1
	if val, err := strconv.Atoi(pageStr); err == nil && val > 0 {
		page = val
	}
	offset := (page - 1) * limit

	// Prepare wildcard strings in Go to avoid SQL syntax confusion
	namePattern := "%" + nameQuery + "%"
	deptPattern := "%" + deptQuery + "%"

	query := fmt.Sprintf(`
		SELECT e.id, e.name, e.email, e.age, e.salary, COALESCE(d.name, 'None') as dept_name
		FROM employees e
		LEFT JOIN departments d ON e.department_id = d.id
		WHERE ($1 = '%%' OR e.name ILIKE $1)
		  AND ($2 = '%%' OR d.name ILIKE $2)
		  AND e.salary BETWEEN $3 AND $4
		ORDER BY %s %s
		LIMIT $5 OFFSET $6
	`, sortColumn, sortDir)

	rows, err := conn.Query(
		context.Background(),
		query,
		namePattern,
		deptPattern,
		minSalary,
		maxSalary,
		limit,
		offset,
	)

	if err != nil {
		fmt.Println("Search error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n--- SEARCH RESULTS ---")
	count := 0
	for rows.Next() {
		var id, age int
		var name, email, deptName string
		var salary float64

		if err := rows.Scan(&id, &name, &email, &age, &salary, &deptName); err != nil {
			continue
		}
		count++
		fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d | Salary: %.2f | Dept: %s\n", id, name, email, age, salary, deptName)
	}

	if count == 0 {
		fmt.Println("No employees found matching criteria.")
	}
}