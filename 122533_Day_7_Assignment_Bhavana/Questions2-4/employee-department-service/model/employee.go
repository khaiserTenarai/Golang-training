package model

import "time"

// Employee mirrors one row in the "new_employee" table.
//
// DepartmentID is a *int (a pointer) instead of a plain int, because
// the column allows NULL - an employee who isn't assigned to any
// department yet. A nil pointer means "no department"; a non-nil
// pointer holds the department's ID.
type Employee struct {
	ID           int
	Name         string
	Email        string
	Age          int
	Salary       float64
	DepartmentID *int
	CreatedAt    time.Time
}
