package dto

import "time"

type CreateDepartmentRequest struct {
	Name           string `json:"name" binding:"required"`
	Code           string `json:"code" binding:"required"`
	HeadcountLimit int    `json:"headcount_limit" binding:"required"`
}

type DepartmentResponse struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Code             string    `json:"code"`
	HeadcountLimit   int       `json:"headcount_limit"`
	CurrentHeadcount int       `json:"current_headcount"`
	CreatedAt        time.Time `json:"created_at"`
}
