package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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
		fmt.Println("\n==================================")
		fmt.Println("  ATTENDANCE & LEAVE MANAGEMENT")
		fmt.Println("==================================")
		fmt.Println("1. Check-In")
		fmt.Println("2. Check-Out")
		fmt.Println("3. Attendance Report")
		fmt.Println("4. Apply for Leave")
		fmt.Println("5. Process Leave (Approve/Reject)")
		fmt.Println("6. View Leave Applications")
		fmt.Println("7. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}

		switch choice {
		case 1:
			checkIn(conn)
		case 2:
			checkOut(conn)
		case 3:
			attendanceReport(conn)
		case 4:
			applyLeave(conn)
		case 5:
			processLeave(conn)
		case 6:
			viewLeaves(conn)
		case 7:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func checkIn(conn *pgx.Conn) {
	fmt.Println("\n----- CHECK-IN -----")
	empIDStr := readInput("Enter Employee ID: ")
	empID, err := strconv.Atoi(empIDStr)
	if err != nil {
		fmt.Println("Invalid Employee ID.")
		return
	}

	_, err = conn.Exec(context.Background(), `INSERT INTO attendance (employee_id) VALUES ($1)`, empID)
	if err != nil {
		fmt.Println("Check-in failed. Ensure Employee ID exists.")
		return
	}
	fmt.Printf("Employee %d checked in successfully at %s!\n", empID, time.Now().Format("15:04:05"))
}

func checkOut(conn *pgx.Conn) {
	fmt.Println("\n----- CHECK-OUT -----")
	empIDStr := readInput("Enter Employee ID: ")
	empID, err := strconv.Atoi(empIDStr)
	if err != nil {
		fmt.Println("Invalid Employee ID.")
		return
	}

	// Update the most recent check-in that doesn't have a check-out yet
	result, err := conn.Exec(
		context.Background(),
		`UPDATE attendance SET check_out = CURRENT_TIMESTAMP 
		 WHERE employee_id = $1 AND check_out IS NULL`,
		empID,
	)

	if err != nil {
		fmt.Println("Error updating check-out:", err)
		return
	}

	if result.RowsAffected() == 0 {
		fmt.Println("No active check-in found for this employee.")
	} else {
		fmt.Printf("Employee %d checked out successfully at %s!\n", empID, time.Now().Format("15:04:05"))
	}
}

func attendanceReport(conn *pgx.Conn) {
	fmt.Println("\n----- ATTENDANCE REPORT -----")
	empIDStr := readInput("Enter Employee ID (leave blank for all): ")
	
	query := `
		SELECT a.id, e.name, a.check_in, a.check_out 
		FROM attendance a
		JOIN employees e ON a.employee_id = e.id
	`
	var rows pgx.Rows
	var err error
	
	if empIDStr == "" {
		query += ` ORDER BY a.check_in DESC`
		rows, err = conn.Query(context.Background(), query)
	} else {
		empID, _ := strconv.Atoi(empIDStr)
		query += ` WHERE a.employee_id = $1 ORDER BY a.check_in DESC`
		rows, err = conn.Query(context.Background(), query, empID)
	}

	if err != nil {
		fmt.Println("Error fetching report:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\n%-5s | %-15s | %-20s | %-20s\n", "ID", "Name", "Check-In", "Check-Out")
	fmt.Println("---------------------------------------------------------------------")
	
	count := 0
	for rows.Next() {
		var id int
		var name string
		var checkIn time.Time
		var checkOut *time.Time // Using a pointer to handle NULLs safely
		
		if err := rows.Scan(&id, &name, &checkIn, &checkOut); err != nil {
			fmt.Println("Error reading row:", err)
			continue
		}
		count++
		
		checkOutStr := "Active / Not Checked Out"
		if checkOut != nil {
			checkOutStr = checkOut.Format("2006-01-02 15:04:05")
		}
		
		fmt.Printf("%-5d | %-15s | %-20s | %-20s\n", id, name, checkIn.Format("2006-01-02 15:04:05"), checkOutStr)
	}
	if count == 0 {
		fmt.Println("No attendance records found.")
	}
}

func applyLeave(conn *pgx.Conn) {
	fmt.Println("\n----- LEAVE APPLICATION -----")
	empIDStr := readInput("Enter Employee ID: ")
	empID, err := strconv.Atoi(empIDStr)
	if err != nil {
		fmt.Println("Invalid Employee ID.")
		return
	}

	leaveType := readInput("Leave Type (e.g., Sick, Casual, Vacation): ")
	startDate := readInput("Start Date (YYYY-MM-DD): ")
	endDate := readInput("End Date (YYYY-MM-DD): ")

	var id int
	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO leaves (employee_id, leave_type, start_date, end_date) 
		 VALUES ($1, $2, $3, $4) RETURNING id`,
		empID, leaveType, startDate, endDate,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error submitting leave application:", err)
	} else {
		fmt.Printf("Leave applied successfully. Application ID: %d\n", id)
	}
}

func processLeave(conn *pgx.Conn) {
	fmt.Println("\n----- PROCESS LEAVE -----")
	leaveIDStr := readInput("Enter Leave Application ID: ")
	leaveID, err := strconv.Atoi(leaveIDStr)
	if err != nil {
		fmt.Println("Invalid Leave ID.")
		return
	}

	action := readInput("Action (1: Approve, 2: Reject): ")
	status := ""
	if action == "1" {
		status = "Approved"
	} else if action == "2" {
		status = "Rejected"
	} else {
		fmt.Println("Invalid action.")
		return
	}

	result, err := conn.Exec(context.Background(), `UPDATE leaves SET status = $1 WHERE id = $2`, status, leaveID)
	if err != nil {
		fmt.Println("Error updating leave status:", err)
		return
	}

	if result.RowsAffected() == 0 {
		fmt.Println("Leave application not found.")
	} else {
		fmt.Printf("Leave ID %d has been %s.\n", leaveID, status)
	}
}

func viewLeaves(conn *pgx.Conn) {
	fmt.Println("\n----- ALL LEAVE APPLICATIONS -----")
	rows, err := conn.Query(
		context.Background(),
		`SELECT l.id, e.name, l.leave_type, l.start_date, l.end_date, l.status 
		 FROM leaves l JOIN employees e ON l.employee_id = e.id 
		 ORDER BY l.id DESC`,
	)
	if err != nil {
		fmt.Println("Error fetching leaves:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\n%-5s | %-15s | %-10s | %-10s | %-10s | %-10s\n", "ID", "Employee", "Type", "Start", "End", "Status")
	fmt.Println("---------------------------------------------------------------------------------")
	
	count := 0
	for rows.Next() {
		var id int
		var name, lType, status string
		var startDate, endDate time.Time
		
		if err := rows.Scan(&id, &name, &lType, &startDate, &endDate, &status); err != nil {
			fmt.Println("Error reading row:", err)
			continue
		}
		count++
		
		fmt.Printf("%-5d | %-15s | %-10s | %-10s | %-10s | %-10s\n", 
			id, name, lType, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), status)
	}
	
	if count == 0 {
		fmt.Println("No leave applications found.")
	}
}