package models

import "time"

// Employee with Gin binding tags (Task 10: JSON binding, Task 12: form binding)
type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name" form:"name" binding:"required,min=2"`
	Email      string    `json:"email" form:"email" binding:"required,email"`
	Age        int       `json:"age" form:"age" binding:"required,min=18,max=65"`
	Department string    `json:"department" form:"department" binding:"required"`
	Salary     float64   `json:"salary" form:"salary" binding:"min=0"`
	CreatedAt  time.Time `json:"created_at"`
}

// PatchEmployee for PATCH — all fields optional (Task 5)
type PatchEmployee struct {
	Name       *string  `json:"name"`
	Email      *string  `json:"email"`
	Age        *int     `json:"age"`
	Department *string  `json:"department"`
	Salary     *float64 `json:"salary"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
