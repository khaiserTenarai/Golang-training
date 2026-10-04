package routes

import (
	"task13_authentication_system/controller"
	"task13_authentication_system/middleware"
	"task13_authentication_system/service"

	"github.com/gorilla/mux"
)

func SetupRoutes(ctrl *controller.AuthController, authService *service.AuthService) *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/auth/register", ctrl.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", ctrl.Login).Methods("POST")

	// Protected routes (require valid token)
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(authService))

	protected.HandleFunc("/auth/logout", ctrl.Logout).Methods("POST")
	protected.HandleFunc("/auth/profile", ctrl.Profile).Methods("GET")

	// Admin-only routes
	admin := r.PathPrefix("/api/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware(authService))
	admin.Use(middleware.AdminOnly)

	admin.HandleFunc("/users", ctrl.ListUsers).Methods("GET")
	admin.HandleFunc("/users/{id}", ctrl.DeleteUser).Methods("DELETE")
	admin.HandleFunc("/users/{id}/role", ctrl.UpdateRole).Methods("PUT")

	return r
}
