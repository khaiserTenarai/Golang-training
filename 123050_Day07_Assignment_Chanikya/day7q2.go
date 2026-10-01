package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	// PostgreSQL connection
	connString := "postgres://postgres:admin@localhost:5432/gotraining"

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	// ==========================================
	// CREATE DEPARTMENT
	// ==========================================

	fmt.Println("\n----- CREATE DEPARTMENT -----")

	var departmentID int

	err = conn.QueryRow(
		context.Background(),
		`
		INSERT INTO departments_q2
		(name, location)
		VALUES ($1,$2)
		RETURNING id
		`,
		"IT",
		"Bangalore",
	).Scan(&departmentID)

	if err != nil {
		log.Println("Department insert error:", err)
	} else {
		fmt.Println("Department created successfully")
		fmt.Println("Department ID:", departmentID)
	}

	// ==========================================
	// READ DEPARTMENT
	// ==========================================

	fmt.Println("\n----- READ DEPARTMENT -----")

	var deptName string
	var deptLocation string

	err = conn.QueryRow(
		context.Background(),
		`
		SELECT name, location
		FROM departments_q2
		WHERE id=$1
		`,
		departmentID,
	).Scan(
		&deptName,
		&deptLocation,
	)

	if err != nil {

		log.Println("Read department error:", err)

	} else {

		fmt.Println("Department Name:", deptName)
		fmt.Println("Location:", deptLocation)

	}

	// ==========================================
	// CREATE EMPLOYEE WITH DEPARTMENT MAPPING
	// ==========================================

	fmt.Println("\n----- CREATE EMPLOYEE -----")

	var employeeID int

	err = conn.QueryRow(
		context.Background(),
		`
		INSERT INTO employees_q2
		(name,email,age,salary,department_id)
		VALUES($1,$2,$3,$4,$5)
		RETURNING id
		`,
		"Rajesh",
		"rajesh_q2@gmail.com",
		30,
		60000,
		departmentID,
	).Scan(&employeeID)

	if err != nil {

		log.Println("Employee insert error:", err)

	} else {

		fmt.Println("Employee created successfully")
		fmt.Println("Employee ID:", employeeID)

	}

	// ==========================================
	// READ EMPLOYEE WITH DEPARTMENT DETAILS
	// JOIN QUERY
	// ==========================================

	fmt.Println("\n----- EMPLOYEE DETAILS -----")

	var employeeName string
	var email string
	var departmentName string

	err = conn.QueryRow(
		context.Background(),
		`
		SELECT 
			e.name,
			e.email,
			d.name

		FROM employees_q2 e

		JOIN departments_q2 d

		ON e.department_id = d.id

		WHERE e.id=$1
		`,
		employeeID,
	).Scan(
		&employeeName,
		&email,
		&departmentName,
	)

	if err != nil {

		log.Println("Employee read error:", err)

	} else {

		fmt.Println("Employee Name:", employeeName)
		fmt.Println("Email:", email)
		fmt.Println("Department:", departmentName)

	}

	// ==========================================
	// UPDATE DEPARTMENT
	// ==========================================

	fmt.Println("\n----- UPDATE DEPARTMENT -----")

	result, err := conn.Exec(
		context.Background(),
		`
		UPDATE departments_q2
		SET location=$1
		WHERE id=$2
		`,
		"Hyderabad",
		departmentID,
	)

	if err != nil {

		log.Println("Update error:", err)

	} else {

		fmt.Println("Department updated successfully")
		fmt.Println("Rows affected:", result.RowsAffected())

	}

	// ==========================================
	// DELETE EMPLOYEE FIRST
	// ==========================================

	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err = conn.Exec(
		context.Background(),
		`
		DELETE FROM employees_q2
		WHERE id=$1
		`,
		employeeID,
	)

	if err != nil {

		log.Println("Employee delete error:", err)

	} else {

		fmt.Println("Employee deleted successfully")
		fmt.Println("Rows affected:", result.RowsAffected())

	}

	// ==========================================
	// DELETE DEPARTMENT
	// ==========================================

	fmt.Println("\n----- DELETE DEPARTMENT -----")

	result, err = conn.Exec(
		context.Background(),
		`
		DELETE FROM departments_q2
		WHERE id=$1
		`,
		departmentID,
	)

	if err != nil {

		log.Println("Department delete error:", err)

	} else {

		fmt.Println("Department deleted successfully")
		fmt.Println("Rows affected:", result.RowsAffected())

	}

}
