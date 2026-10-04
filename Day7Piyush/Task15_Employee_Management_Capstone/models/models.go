package models

import "time"

// ---- User & Auth ----

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Password     string    `json:"password,omitempty"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	User    *User  `json:"user,omitempty"`
}

// ---- Department ----

type Department struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
}

// ---- Employee ----

type Employee struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Age            int       `json:"age"`
	DepartmentID   *int      `json:"department_id"`
	DepartmentName string    `json:"department_name,omitempty"`
	Salary         float64   `json:"salary"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SalaryHistory struct {
	ID           int       `json:"id"`
	EmployeeID   int       `json:"employee_id"`
	OldSalary    float64   `json:"old_salary"`
	NewSalary    float64   `json:"new_salary"`
	Reason       string    `json:"reason"`
	ChangedAt    time.Time `json:"changed_at"`
}

// ---- Common ----

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
}

type SalaryUpdateRequest struct {
	NewSalary float64 `json:"new_salary"`
	Reason    string  `json:"reason"`
}
