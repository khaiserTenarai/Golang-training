package models

import "time"

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}

type Attendance struct {
	ID           int        `json:"id"`
	EmployeeID   int        `json:"employee_id"`
	EmployeeName string     `json:"employee_name,omitempty"`
	CheckIn      *time.Time `json:"check_in"`
	CheckOut     *time.Time `json:"check_out"`
	Date         string     `json:"date"`
	Status       string     `json:"status"`
}

type LeaveRequest struct {
	ID           int       `json:"id"`
	EmployeeID   int       `json:"employee_id"`
	EmployeeName string    `json:"employee_name,omitempty"`
	LeaveType    string    `json:"leave_type"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
	Reason       string    `json:"reason"`
	Status       string    `json:"status"`
	ReviewedBy   string    `json:"reviewed_by,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
