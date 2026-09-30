package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var allowedSortColumns = map[string]bool{
	"name":       true,
	"department": true,
	"salary":     true,
}

func main() {
	connString := "postgres://postgres:password@localhost:5432/searchdb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	fmt.Print("Name contains (blank to skip): ")
	name := readLine(reader)

	fmt.Print("Department (blank to skip): ")
	department := readLine(reader)

	fmt.Print("Minimum salary (blank to skip): ")
	minSalaryStr := readLine(reader)

	fmt.Print("Sort by (name/department/salary): ")
	sortBy := readLine(reader)
	if !allowedSortColumns[sortBy] {
		sortBy = "name"
	}

	fmt.Print("Sort direction (asc/desc): ")
	direction := readLine(reader)
	if direction != "asc" && direction != "desc" {
		direction = "asc"
	}

	fmt.Print("Page number (starting at 1): ")
	page, _ := strconv.Atoi(readLine(reader))
	if page < 1 {
		page = 1
	}

	fmt.Print("Page size: ")
	pageSize, _ := strconv.Atoi(readLine(reader))
	if pageSize < 1 {
		pageSize = 10
	}

	query := "SELECT id, name, department, salary FROM employees WHERE 1=1"
	args := []interface{}{}
	argPos := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argPos)
		args = append(args, "%"+name+"%")
		argPos++
	}
	if department != "" {
		query += fmt.Sprintf(" AND department = $%d", argPos)
		args = append(args, department)
		argPos++
	}
	if minSalaryStr != "" {
		minSalary, _ := strconv.ParseFloat(minSalaryStr, 64)
		query += fmt.Sprintf(" AND salary >= $%d", argPos)
		args = append(args, minSalary)
		argPos++
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortBy, direction, argPos, argPos+1)
	args = append(args, pageSize, (page-1)*pageSize)

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var id int
		var empName, empDept string
		var salary float64
		if err := rows.Scan(&id, &empName, &empDept, &salary); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Name: %s | Department: %s | Salary: %.2f\n", id, empName, empDept, salary)
		found = true
	}
	if !found {
		fmt.Println("No matching employees")
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
