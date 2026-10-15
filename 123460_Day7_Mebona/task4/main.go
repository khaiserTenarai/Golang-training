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
		fmt.Println("      SALARY MANAGEMENT")
		fmt.Println("==============================")
		fmt.Println("1. Update Employee Salary & Track History")
		fmt.Println("2. View Salary History for Employee")
		fmt.Println("3. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 3.")
			continue
		}

		switch choice {
		case 1:
			updateSalaryWithTransaction(conn)
		case 2:
			viewSalaryHistory(conn)
		case 3:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 3.")
		}
	}
}

func updateSalaryWithTransaction(conn *pgx.Conn) {
	fmt.Println("\n----- UPDATE SALARY -----")
	
	idStr := readInput("Enter Employee ID: ")
	empID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid Employee ID.")
		return
	}

	newSalaryStr := readInput("Enter New Salary: ")
	newSalary, err := strconv.ParseFloat(newSalaryStr, 64)
	if err != nil {
		fmt.Println("Invalid salary amount.")
		return
	}

	ctx := context.Background()

	// Start a PostgreSQL transaction using pgx
	tx, err := conn.Begin(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction:", err)
		return
	}
	// Defer rollback to ensure it aborts if not committed successfully
	defer tx.Rollback(ctx)

	// Fetch current salary inside the transaction
	var currentSalary float64
	err = tx.QueryRow(ctx, `SELECT salary FROM employees WHERE id = $1`, empID).Scan(&currentSalary)
	if err != nil {
		fmt.Println("Error: Employee not found or database error:", err)
		return
	}

	if currentSalary == newSalary {
		fmt.Println("The new salary is identical to the current salary. No update needed.")
		return
	}

	// Update employee's salary
	_, err = tx.Exec(ctx, `UPDATE employees SET salary = $1 WHERE id = $2`, newSalary, empID)
	if err != nil {
		fmt.Println("Failed to update employee salary:", err)
		return
	}

	// Record history in salary_history table
	_, err = tx.Exec(
		ctx,
		`INSERT INTO salary_history (employee_id, old_salary, new_salary) VALUES ($1, $2, $3)`,
		empID, currentSalary, newSalary,
	)
	if err != nil {
		fmt.Println("Failed to record salary history:", err)
		return
	}

	// Commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}

	fmt.Printf("Salary updated successfully for Employee ID %d (Old: %.2f -> New: %.2f)\n", empID, currentSalary, newSalary)
}

func viewSalaryHistory(conn *pgx.Conn) {
	fmt.Println("\n----- VIEW SALARY HISTORY -----")
	idStr := readInput("Enter Employee ID to view history: ")
	empID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid Employee ID.")
		return
	}

	rows, err := conn.Query(
		context.Background(),
		`SELECT h.id, e.name, h.old_salary, h.new_salary, h.changed_at 
		 FROM salary_history h
		 JOIN employees e ON h.employee_id = e.id
		 WHERE h.employee_id = $1
		 ORDER BY h.changed_at DESC`,
		empID,
	)
	if err != nil {
		fmt.Println("Error fetching history:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\n--- Salary History for Employee ID: %d ---\n", empID)
	count := 0
	for rows.Next() {
		var historyID int
		var name string
		var oldSal, newSal float64
		var changedAt string

		if err := rows.Scan(&historyID, &name, &oldSal, &newSal, &changedAt); err != nil {
			continue
		}
		count++
		fmt.Printf("Name: %s | Old Salary: %.2f -> New Salary: %.2f | Date: %s\n", name, oldSal, newSal, changedAt)
	}

	if count == 0 {
		fmt.Println("No salary history found for this employee.")
	}
}