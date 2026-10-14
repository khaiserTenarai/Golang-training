package routes

import (
	"net/http"

	"employee-rest-api/controller"
)

func Register(mux *http.ServeMux, c *controller.EmployeeController) {
	// Go's standard ServeMux supports method-aware patterns in Go 1.22+.
	mux.HandleFunc("POST /api/employees", c.Create)
	mux.HandleFunc("GET /api/employees", c.GetAll)
	mux.HandleFunc("GET /api/employees/{id}", c.GetByID)
	mux.HandleFunc("PUT /api/employees/{id}", c.Update)
	mux.HandleFunc("DELETE /api/employees/{id}", c.Delete)
}
