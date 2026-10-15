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
		fmt.Println("    DEPARTMENT MANAGEMENT")
		fmt.Println("==============================")
		fmt.Println("1. Create Department")
		fmt.Println("2. List All Departments")
		fmt.Println("3. Update Department")
		fmt.Println("4. Delete Department")
		fmt.Println("5. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			createDepartment(conn)
		case 2:
			listDepartments(conn)
		case 3:
			updateDepartment(conn)
		case 4:
			deleteDepartment(conn)
		case 5:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 5.")
		}
	}
}

func createDepartment(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE DEPARTMENT -----")
	name := readInput("Department Name: ")
	location := readInput("Location: ")

	var id int
	err := conn.QueryRow(
		context.Background(),
		`INSERT INTO departments (name, location) VALUES ($1, $2) RETURNING id`,
		name, location,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error creating department:", err)
	} else {
		fmt.Printf("Department created successfully with ID: %d\n", id)
	}
}

func listDepartments(conn *pgx.Conn) {
	fmt.Println("\n----- ALL DEPARTMENTS -----")
	rows, err := conn.Query(context.Background(), `SELECT id, name, location FROM departments ORDER BY id`)
	if err != nil {
		fmt.Println("Error reading departments:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, location string
		if err := rows.Scan(&id, &name, &location); err != nil {
			continue
		}
		fmt.Printf("ID: %d | Name: %s | Location: %s\n", id, name, location)
	}
}

func updateDepartment(conn *pgx.Conn) {
	fmt.Println("\n----- UPDATE DEPARTMENT -----")
	idStr := readInput("Enter Department ID to update: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	name := readInput("New Department Name: ")
	location := readInput("New Location: ")

	result, err := conn.Exec(
		context.Background(),
		`UPDATE departments SET name = $1, location = $2 WHERE id = $3`,
		name, location, id,
	)
	if err != nil {
		fmt.Println("Error updating department:", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("No department found with that ID.")
	} else {
		fmt.Println("Department updated successfully.")
	}
}

func deleteDepartment(conn *pgx.Conn) {
	fmt.Println("\n----- DELETE DEPARTMENT -----")
	idStr := readInput("Enter Department ID to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	result, err := conn.Exec(context.Background(), `DELETE FROM departments WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error deleting department (make sure no employees are referencing it):", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("No department found with that ID.")
	} else {
		fmt.Println("Department deleted successfully.")
	}
}