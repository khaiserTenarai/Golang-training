package model

import "time"

type Account struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	HolderName   string    `json:"holder_name"`
	Balance      float64   `json:"balance"`
	DocumentPath string    `json:"document_path,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateAccountRequest struct {
	HolderName     string  `json:"holder_name" binding:"required,min=2,max=100"`
	InitialDeposit float64 `json:"initial_deposit" binding:"gte=0"`
}


type PageQuery struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

func (q *PageQuery) Normalize() {
	if q.Page == 0 {
		q.Page = 1
	}
	if q.Limit == 0 {
		q.Limit = 10
	}
}

func (q PageQuery) Offset() int { return (q.Page - 1) * q.Limit }

type PageResult[T any] struct {
	Items      []T `json:"items"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func NewPageResult[T any](items []T, q PageQuery, total int) PageResult[T] {
	if items == nil {
		items = []T{}
	}
	pages := (total + q.Limit - 1) / q.Limit
	return PageResult[T]{Items: items, Page: q.Page, Limit: q.Limit, Total: total, TotalPages: pages}
}
