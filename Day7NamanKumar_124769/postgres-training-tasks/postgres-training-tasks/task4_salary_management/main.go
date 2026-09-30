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

func main() {
	connString := "postgres://postgres:password@localhost:5432/salarydb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Employee")
		fmt.Println("2. Update Salary")
		fmt.Println("3. View Salary History")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createEmployee(ctx, db, reader)
		case "2":
			updateSalary(ctx, db, reader)
		case "3":
			viewHistory(ctx, db, reader)
		case "4":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func createEmployee(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Name: ")
	name := readLine(reader)

	fmt.Print("Starting salary: ")
	salary, _ := strconv.ParseFloat(readLine(reader), 64)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO employees (name, salary) VALUES ($1, $2) RETURNING id`, name, salary).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created employee with ID:", id)
}

func updateSalary(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("New salary: ")
	newSalary, _ := strconv.ParseFloat(readLine(reader), 64)

	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer tx.Rollback(ctx)

	var oldSalary float64
	err = tx.QueryRow(ctx, `SELECT salary FROM employees WHERE id = $1 FOR UPDATE`, id).Scan(&oldSalary)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = tx.Exec(ctx, `UPDATE employees SET salary = $1 WHERE id = $2`, newSalary, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = tx.Exec(ctx, `INSERT INTO salary_history (employee_id, old_salary, new_salary) VALUES ($1, $2, $3)`, id, oldSalary, newSalary)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Salary updated and history recorded")
}

func viewHistory(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	rows, err := db.Query(ctx, `SELECT old_salary, new_salary, changed_at FROM salary_history WHERE employee_id = $1 ORDER BY changed_at`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var oldSalary, newSalary float64
		var changedAt string
		if err := rows.Scan(&oldSalary, &newSalary, &changedAt); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%s: %.2f -> %.2f\n", changedAt, oldSalary, newSalary)
		found = true
	}
	if !found {
		fmt.Println("No salary history found")
	}
}
