package models

import "github.com/golang-jwt/jwt/v5"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=3"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

type Claims struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
	jwt.RegisteredClaims
}

type PaginatedProductsResponse struct {
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	Total int       `json:"total"`
	Data  []Product `json:"data"`
}
