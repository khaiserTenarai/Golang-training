package routes

import (
	"task12_repository_service_architecture/controller"

	"github.com/gorilla/mux"
)

func SetupRoutes(ctrl *controller.EmployeeController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/employees", ctrl.Create).Methods("POST")
	r.HandleFunc("/api/employees", ctrl.GetAll).Methods("GET")
	r.HandleFunc("/api/employees/{id}", ctrl.GetByID).Methods("GET")
	r.HandleFunc("/api/employees/{id}", ctrl.Update).Methods("PUT")
	r.HandleFunc("/api/employees/{id}", ctrl.Delete).Methods("DELETE")

	return r
}
