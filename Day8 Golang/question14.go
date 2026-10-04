package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Attendance struct {
	ID       int
	EmpID    int
	CheckIn  time.Time
	CheckOut sql.NullTime
	WorkDate time.Time
}

type Leave struct {
	ID        int
	EmpID     int
	StartDate time.Time
	EndDate   time.Time
	Reason    string
	Status    string
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=Day7 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var empID int
	err = db.QueryRow(`INSERT INTO employees (name) VALUES ('John Doe') RETURNING id`).Scan(&empID)
	if err != nil {
		log.Fatal(err)
	}

	checkIn(db, empID)
	
	time.Sleep(2 * time.Second) 
	
	checkOut(db, empID)

	report := getAttendanceReport(db, empID, time.Now().AddDate(0, 0, -7), time.Now())
	for _, a := range report {
		fmt.Printf("Work Date: %s | In: %s | Out: %v\n", 
			a.WorkDate.Format("2006-01-02"), 
			a.CheckIn.Format("15:04:05"), 
			a.CheckOut.Time.Format("15:04:05"))
	}

	leaveID := applyLeave(db, empID, "2026-10-10", "2026-10-15", "Vacation")
	
	updateLeaveStatus(db, leaveID, "approved")
}

func checkIn(db *sql.DB, empID int) {
	_, err := db.Exec(`
		INSERT INTO attendance (employee_id, check_in, work_date) 
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_DATE)`, empID)
	if err != nil {
		log.Fatal(err)
	}
}

func checkOut(db *sql.DB, empID int) {
	_, err := db.Exec(`
		UPDATE attendance 
		SET check_out = CURRENT_TIMESTAMP 
		WHERE employee_id = $1 AND work_date = CURRENT_DATE AND check_out IS NULL`, empID)
	if err != nil {
		log.Fatal(err)
	}
}

func getAttendanceReport(db *sql.DB, empID int, startDate, endDate time.Time) []Attendance {
	rows, err := db.Query(`
		SELECT id, employee_id, check_in, check_out, work_date 
		FROM attendance 
		WHERE employee_id = $1 AND work_date >= $2 AND work_date <= $3`, 
		empID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var records []Attendance
	for rows.Next() {
		var a Attendance
		if err := rows.Scan(&a.ID, &a.EmpID, &a.CheckIn, &a.CheckOut, &a.WorkDate); err != nil {
			log.Fatal(err)
		}
		records = append(records, a)
	}
	return records
}

func applyLeave(db *sql.DB, empID int, startDate, endDate, reason string) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO leaves (employee_id, start_date, end_date, reason) 
		VALUES ($1, $2, $3, $4) RETURNING id`, 
		empID, startDate, endDate, reason).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func updateLeaveStatus(db *sql.DB, leaveID int, status string) {
	_, err := db.Exec(`UPDATE leaves SET status = $1 WHERE id = $2`, status, leaveID)
	if err != nil {
		log.Fatal(err)
	}
}