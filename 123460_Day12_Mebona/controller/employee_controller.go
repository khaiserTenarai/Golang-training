package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"employee-management/config"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/service"
)

type EmployeeController struct {
	service   service.EmployeeService
	jwtSecret string
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	cfg := config.Load()
	return &EmployeeController{service: s, jwtSecret: cfg.JWTSecret}
}

func (c *EmployeeController) RegisterRoutes(mux *http.ServeMux) {
	// Standard Liveness/Readiness Probes
	mux.HandleFunc("/health", c.Health)
	mux.HandleFunc("/ready", c.Ready)

	// Helper route to issue JWT token for testing in SoapUI
	mux.HandleFunc("/token", c.GetToken)

	// Middleware chain for protected routes
	authMiddleware := middleware.Auth(c.jwtSecret)

	// Employees router multiplexer with HTTP method filtering
	empHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			c.GetAll(w, r)
		case http.MethodPost:
			middleware.AuthorizeRole("admin", c.Create)(w, r)
		default:
			http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	// Register authenticated route path
	mux.Handle("/api/v1/employees", authMiddleware(empHandler))
}

func (c *EmployeeController) GetToken(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	if role == "" {
		role = "admin"
	}
	token, err := middleware.GenerateToken("testuser", role, c.jwtSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, `{"error":"invalid request payload"}`, http.StatusBadRequest)
		return
	}

	if err := c.service.CreateEmployee(r.Context(), &emp); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	filter := model.EmployeeFilter{
		Department: query.Get("department"),
		SortBy:     query.Get("sortBy"),
		Order:      query.Get("order"),
		Page:       page,
		Limit:      limit,
	}

	employees, total, err := c.service.GetEmployees(r.Context(), filter)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch employees"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data":  employees,
		"total": total,
		"page":  filter.Page,
		"limit": filter.Limit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *EmployeeController) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}

func (c *EmployeeController) Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := c.service.CheckHealth(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"status": "DOWN", "database": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "READY", "database": "CONNECTED"})
}