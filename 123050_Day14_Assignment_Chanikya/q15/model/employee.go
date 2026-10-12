package model

import "time"

type Employee struct {
 ID int64 `json:"id"`
 Name string `json:"name"`
 Email string `json:"email"`
 Age int `json:"age"`
 Salary float64 `json:"salary"`
 CreatedAt time.Time `json:"created_at,omitempty"`
 UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type EmployeeInput struct {
 Name string `json:"name"`
 Email string `json:"email"`
 Age int `json:"age"`
 Salary float64 `json:"salary"`
}
type EmployeePatch struct {
 Name *string `json:"name"`
 Email *string `json:"email"`
 Age *int `json:"age"`
 Salary *float64 `json:"salary"`
}
