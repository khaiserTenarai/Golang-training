// Package controller contains HTTP request handlers.
package controller

// Import standard packages required by the REST controller.
import (
	// Import JSON encoding and decoding.
	"encoding/json"
	// Import errors for database error comparison.
	"errors"
	// Import net/http for HTTP handling.
	"net/http"
	// Import strconv for path parameter conversion.
	"strconv"
	// Import strings for URL path processing.
	"strings"
)

// Import pgx for PostgreSQL error comparison.
import "github.com/jackc/pgx/v5"

// Import employee model.
import "employee-management/model"

// Import employee service.
import "employee-management/service"

// EmployeeController handles employee HTTP requests.
type EmployeeController struct {
	// service stores the business service dependency.
	service service.EmployeeService
}

// NewEmployeeController creates a controller with dependency injection.
func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	// Return the controller instance.
	return &EmployeeController{service: service}
}

// RegisterRoutes registers all employee REST endpoints.
func (c *EmployeeController) RegisterRoutes(mux *http.ServeMux) {
	// Register the collection endpoint for GET and POST.
	mux.HandleFunc("/employees", c.handleCollection)
	// Register the resource endpoint for GET, PUT and DELETE with a path parameter.
	mux.HandleFunc("/employees/", c.handleByID)
}

// handleCollection handles collection-level employee requests.
func (c *EmployeeController) handleCollection(w http.ResponseWriter, r *http.Request) {
	// Handle GET requests for all employees.
	if r.Method == http.MethodGet {}
		// Read the optional sort query parameter.
		sortBy := r.URL.Query().Get("sort")
		// Call the service to retrieve employees.
		employees, err := c.service.GetAll(r.Context(), sortBy)
		// Handle service errors.
		if err != nil {
			// Return an internal server error response.
			writeError(w, http.StatusInternalServerError, err.Error())
			// Stop request processing.
			return
		}
		// Return employees as JSON.
		writeJSON(w, http.StatusOK, employees)
		// Stop request processing.
		return
	}
	// Handle POST requests for creating an employee.
	if r.Method == http.MethodPost {
		// Create an empty employee model.
		var employee model.Employee
		// Decode JSON request body.
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			// Return a bad request response for invalid JSON.
			writeError(w, http.StatusBadRequest, "invalid JSON request body")
			// Stop request processing.
			return
		}
		// Call the service to create the employee.
		if err := c.service.Create(r.Context(), &employee); err != nil {
			// Return validation errors as bad requests.
			writeError(w, http.StatusBadRequest, err.Error())
			// Stop request processing.
			return
		}
		// Return the newly created employee.
		writeJSON(w, http.StatusCreated, employee)
		// Stop request processing.
		return
	}
	// Return method-not-allowed for unsupported HTTP methods.
	writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// handleByID handles employee resource requests containing a path parameter.
func (c *EmployeeController) handleByID(w http.ResponseWriter, r *http.Request) {
	// Remove the /employees/ prefix from the URL.
	idText := strings.TrimPrefix(r.URL.Path, "/employees/")
	// Reject empty IDs.
	if idText == "" {
		// Return a bad request response.
		writeError(w, http.StatusBadRequest, "employee id is required")
		// Stop request processing.
		return
	}
	// Convert the path parameter to an integer.
	id, err := strconv.ParseInt(idText, 10, 64)
	// Reject invalid numeric IDs.
	if err != nil {
		// Return a bad request response.
		writeError(w, http.StatusBadRequest, "employee id must be numeric")
		// Stop request processing.
		return
	}
	// Handle GET for a single employee.
	if r.Method == http.MethodGet {
		// Retrieve the employee using the service.
		employee, err := c.service.GetByID(r.Context(), id)
		// Handle missing employee records.
		if errors.Is(err, pgx.ErrNoRows) {
			// Return HTTP 404.
			writeError(w, http.StatusNotFound, "employee not found")
			// Stop request processing.
			return
		}
		// Handle other service or database errors.
		if err != nil {
			// Return HTTP 500.
			writeError(w, http.StatusInternalServerError, err.Error())
			// Stop request processing.
			return
		}
		// Return the employee as JSON.
		writeJSON(w, http.StatusOK, employee)
		// Stop request processing.
		return
	}
	// Handle PUT for updating an employee.
	if r.Method == http.MethodPut {
		// Create an employee model for the JSON request.
		var employee model.Employee
		// Decode the JSON request body.
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			// Return a bad request for invalid JSON.
			writeError(w, http.StatusBadRequest, "invalid JSON request body")
			// Stop request processing.
			return
		}
		// Update the employee using the service.
		updated, err := c.service.Update(r.Context(), id, &employee)
		// Handle missing employees.
		if errors.Is(err, pgx.ErrNoRows) {
			// Return HTTP 404.
			writeError(w, http.StatusNotFound, "employee not found")
			// Stop request processing.
			return
		}
		// Handle validation errors.
		if err != nil {
			// Return HTTP 400 for update validation failures.
			writeError(w, http.StatusBadRequest, err.Error())
			// Stop request processing.
			return
		}
		// Return the updated employee.
		writeJSON(w, http.StatusOK, updated)
		// Stop request processing.
		return
	}
	// Handle DELETE for an employee.
	if r.Method == http.MethodDelete {
		// Delete the employee using the service.
		err := c.service.Delete(r.Context(), id)
		// Handle a missing employee.
		if err != nil && err.Error() == "employee not found" {
			// Return HTTP 404.
			writeError(w, http.StatusNotFound, "employee not found")
			// Stop request processing.
			return
		}
		// Handle other delete errors.
		if err != nil {
			// Return HTTP 500.
			writeError(w, http.StatusInternalServerError, err.Error())
			// Stop request processing.
			return
		}
		// Return a successful empty response.
		w.WriteHeader(http.StatusNoContent)
		// Stop request processing.
		return
	}
	// Return method-not-allowed for unsupported methods.
	writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// writeJSON writes an HTTP JSON response.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	// Set the JSON content type.
	w.Header().Set("Content-Type", "application/json")
	// Write the HTTP status.
	w.WriteHeader(status)
	// Encode the response as JSON.
	_ = json.NewEncoder(w).Encode(data)
}

// writeError writes a consistent JSON error response.
func writeError(w http.ResponseWriter, status int, message string) {
	// Create a simple error response object.
	response := map[string]string{"error": message}
	// Write the error as JSON.
	writeJSON(w, status, response)
}
