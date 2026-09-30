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
	connString := "postgres://postgres:password@localhost:5432/attendancedb"
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
		fmt.Println("2. Check In")
		fmt.Println("3. Check Out")
		fmt.Println("4. Attendance Report")
		fmt.Println("5. Apply for Leave")
		fmt.Println("6. Approve Leave")
		fmt.Println("7. Reject Leave")
		fmt.Println("8. List Leave Requests")
		fmt.Println("9. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createEmployee(ctx, db, reader)
		case "2":
			checkIn(ctx, db, reader)
		case "3":
			checkOut(ctx, db, reader)
		case "4":
			attendanceReport(ctx, db, reader)
		case "5":
			applyLeave(ctx, db, reader)
		case "6":
			updateLeaveStatus(ctx, db, reader, "approved")
		case "7":
			updateLeaveStatus(ctx, db, reader, "rejected")
		case "8":
			listLeaveRequests(ctx, db)
		case "9":
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

	var id int
	err := db.QueryRow(ctx, `INSERT INTO employees (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created employee with ID:", id)
}

func checkIn(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	var attendanceID int
	err := db.QueryRow(ctx, `INSERT INTO attendance (employee_id, check_in) VALUES ($1, NOW()) RETURNING id`, id).Scan(&attendanceID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Checked in, attendance ID:", attendanceID)
}

func checkOut(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Attendance ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	result, err := db.Exec(ctx, `UPDATE attendance SET check_out = NOW() WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Attendance record not found")
		return
	}
	fmt.Println("Checked out")
}

func attendanceReport(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	rows, err := db.Query(ctx, `SELECT check_in, check_out FROM attendance WHERE employee_id = $1 ORDER BY check_in`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var checkIn, checkOut *string
		if err := rows.Scan(&checkIn, &checkOut); err != nil {
			fmt.Println("Error:", err)
			return
		}
		outStr := "still checked in"
		if checkOut != nil {
			outStr = *checkOut
		}
		fmt.Printf("Check-in: %s | Check-out: %s\n", *checkIn, outStr)
		found = true
	}
	if !found {
		fmt.Println("No attendance records found")
	}
}

func applyLeave(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Start date (YYYY-MM-DD): ")
	startDate := readLine(reader)

	fmt.Print("End date (YYYY-MM-DD): ")
	endDate := readLine(reader)

	var leaveID int
	err := db.QueryRow(ctx, `INSERT INTO leave_requests (employee_id, start_date, end_date) VALUES ($1, $2, $3) RETURNING id`,
		id, startDate, endDate).Scan(&leaveID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Leave request submitted with ID:", leaveID)
}

func updateLeaveStatus(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader, status string) {
	fmt.Print("Leave request ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	result, err := db.Exec(ctx, `UPDATE leave_requests SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Leave request not found")
		return
	}
	fmt.Println("Leave request", status)
}

func listLeaveRequests(ctx context.Context, db *pgxpool.Pool) {
	rows, err := db.Query(ctx, `SELECT id, employee_id, start_date, end_date, status FROM leave_requests ORDER BY id`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, empID int
		var startDate, endDate, status string
		if err := rows.Scan(&id, &empID, &startDate, &endDate, &status); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Employee: %d | %s to %s | Status: %s\n", id, empID, startDate, endDate, status)
	}
}
