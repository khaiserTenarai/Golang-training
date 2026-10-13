package router

import (
	"net/http"

	"employee-api-go/internal/auth"
	"employee-api-go/internal/handler"
	"employee-api-go/internal/middleware"
)

func New(employees *handler.EmployeeHandler, authHandler *handler.AuthHandler, system *handler.SystemHandler, tokens *auth.TokenManager) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", system.Health)
	mux.HandleFunc("GET /ready", system.Ready)
	mux.HandleFunc("POST /auth/token", authHandler.Token)

	mux.Handle("POST /api/employees", protectedAdmin(tokens, http.HandlerFunc(employees.Create)))
	mux.Handle("GET /api/employees", protectedUser(tokens, http.HandlerFunc(employees.List)))
	mux.Handle("GET /api/employees/{id}", protectedUser(tokens, http.HandlerFunc(employees.GetByID)))
	mux.Handle("PUT /api/employees/{id}", protectedAdmin(tokens, http.HandlerFunc(employees.Update)))
	mux.Handle("DELETE /api/employees/{id}", protectedAdmin(tokens, http.HandlerFunc(employees.Delete)))
	return mux
}

func protectedUser(tokens *auth.TokenManager, next http.Handler) http.Handler {
	return middleware.Authentication(tokens, next)
}

func protectedAdmin(tokens *auth.TokenManager, next http.Handler) http.Handler {
	return middleware.Authentication(tokens, middleware.RequireRole("admin", next))
}
