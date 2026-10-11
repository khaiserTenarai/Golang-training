package routes

import (
	"task14_15_employee_api/controller"
	"task14_15_employee_api/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(ctrl *controller.EmployeeController) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware) // Task 12

	r.HandleFunc("/api/employees", ctrl.GetAll).Methods("GET")      // Tasks 2, 11
	r.HandleFunc("/api/employees/{id}", ctrl.GetByID).Methods("GET") // Tasks 3, 10
	r.HandleFunc("/api/employees", ctrl.Create).Methods("POST")      // Tasks 4, 9
	r.HandleFunc("/api/employees/{id}", ctrl.Update).Methods("PUT")  // Task 5
	r.HandleFunc("/api/employees/{id}", ctrl.Delete).Methods("DELETE") // Task 6

	return r
}
