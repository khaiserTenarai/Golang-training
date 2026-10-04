package routes

import (
	"task11_employee_rest_api/controller"

	"github.com/gorilla/mux"
)

func SetupRoutes(ctrl *controller.EmployeeController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/employees", ctrl.CreateEmployee).Methods("POST")
	r.HandleFunc("/api/employees", ctrl.GetAllEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", ctrl.GetEmployee).Methods("GET")
	r.HandleFunc("/api/employees/{id}", ctrl.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees/{id}", ctrl.DeleteEmployee).Methods("DELETE")

	return r
}
