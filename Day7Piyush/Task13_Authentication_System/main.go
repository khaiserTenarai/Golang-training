package main

import (
	"fmt"
	"log"
	"net/http"
	"task13_authentication_system/config"
	"task13_authentication_system/controller"
	"task13_authentication_system/repository"
	"task13_authentication_system/routes"
	"task13_authentication_system/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(repo)
	ctrl := controller.NewAuthController(authService)
	router := routes.SetupRoutes(ctrl, authService)

	fmt.Println("Authentication Server running on http://localhost:8082")
	fmt.Println("Public Endpoints:")
	fmt.Println("  POST /api/auth/register  - Register new user")
	fmt.Println("  POST /api/auth/login     - Login (returns token)")
	fmt.Println("Protected Endpoints (Bearer token required):")
	fmt.Println("  POST /api/auth/logout    - Logout")
	fmt.Println("  GET  /api/auth/profile   - View profile")
	fmt.Println("Admin Endpoints (admin role required):")
	fmt.Println("  GET    /api/admin/users        - List all users")
	fmt.Println("  DELETE /api/admin/users/{id}    - Delete user")
	fmt.Println("  PUT    /api/admin/users/{id}/role - Update role")
	log.Fatal(http.ListenAndServe(":8082", router))
}
