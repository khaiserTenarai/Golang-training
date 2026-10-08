package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:sasi2356@localhost:5432/assignment_db"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	// Schema setup
	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS attendance (
			id SERIAL PRIMARY KEY,
			employee_id INT NOT NULL,
			check_in TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			check_out TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS leave_requests (
			id SERIAL PRIMARY KEY,
			employee_id INT NOT NULL,
			reason TEXT NOT NULL,
			status VARCHAR(20) DEFAULT 'PENDING'
		);
	`)

	// 1. Check-in Employee #1
	var attID int
	_ = conn.QueryRow(ctx, "INSERT INTO attendance (employee_id) VALUES ($1) RETURNING id", 1).Scan(&attID)
	fmt.Printf("[ATTENDANCE] Employee #1 checked in with record ID: %d\n", attID)

	// 2. Check-out Employee #1
	_, _ = conn.Exec(ctx, "UPDATE attendance SET check_out = CURRENT_TIMESTAMP WHERE id = $1", attID)
	fmt.Printf("[ATTENDANCE] Employee #1 checked out successfully.\n")

	// 3. Apply for Leave
	var leaveID int
	_ = conn.QueryRow(ctx, "INSERT INTO leave_requests (employee_id, reason) VALUES ($1, $2) RETURNING id", 1, "Medical Leave").Scan(&leaveID)
	fmt.Printf("[LEAVE] Leave Request #%d created with status: PENDING\n", leaveID)

	// 4. Approve Leave
	_, _ = conn.Exec(ctx, "UPDATE leave_requests SET status = 'APPROVED' WHERE id = $1", leaveID)

	// Query status
	var status string
	_ = conn.QueryRow(ctx, "SELECT status FROM leave_requests WHERE id = $1", leaveID).Scan(&status)
	fmt.Printf("[LEAVE] Leave Request #%d updated status: %s\n", leaveID, status)
}