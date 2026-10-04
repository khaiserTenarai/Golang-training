package routes

import (
	"task15_employee_management_capstone/controller"
	"task15_employee_management_capstone/middleware"
	"task15_employee_management_capstone/service"

	"github.com/gorilla/mux"
)

func SetupRoutes(
	authCtrl *controller.AuthController,
	empCtrl *controller.EmployeeController,
	deptCtrl *controller.DepartmentController,
	authSvc service.AuthService,
) *mux.Router {
	r := mux.NewRouter()

	// Global logging middleware
	r.Use(middleware.LoggingMiddleware)

	// ---- Public routes ----
	r.HandleFunc("/api/auth/register", authCtrl.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", authCtrl.Login).Methods("POST")

	// ---- Protected routes (require auth token) ----
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(authSvc))

	// Auth
	protected.HandleFunc("/auth/logout", authCtrl.Logout).Methods("POST")
	protected.HandleFunc("/auth/profile", authCtrl.Profile).Methods("GET")

	// Employees (CRUD + pagination + salary transactions)
	protected.HandleFunc("/employees", empCtrl.Create).Methods("POST")
	protected.HandleFunc("/employees", empCtrl.GetAll).Methods("GET")
	protected.HandleFunc("/employees/{id}", empCtrl.GetByID).Methods("GET")
	protected.HandleFunc("/employees/{id}", empCtrl.Update).Methods("PUT")
	protected.HandleFunc("/employees/{id}", empCtrl.Delete).Methods("DELETE")
	protected.HandleFunc("/employees/{id}/salary", empCtrl.UpdateSalary).Methods("PUT")
	protected.HandleFunc("/employees/{id}/salary-history", empCtrl.SalaryHistory).Methods("GET")

	// Departments
	protected.HandleFunc("/departments", deptCtrl.Create).Methods("POST")
	protected.HandleFunc("/departments", deptCtrl.GetAll).Methods("GET")
	protected.HandleFunc("/departments/{id}", deptCtrl.GetByID).Methods("GET")
	protected.HandleFunc("/departments/{id}", deptCtrl.Update).Methods("PUT")
	protected.HandleFunc("/departments/{id}", deptCtrl.Delete).Methods("DELETE")

	return r
}
