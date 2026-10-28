package model

type Employee struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"gte=0"`
}

type EmployeeListResponse struct {
	Data       []Employee `json:"data"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	Total      int64      `json:"total"`
	TotalPages int        `json:"total_pages"`
}
