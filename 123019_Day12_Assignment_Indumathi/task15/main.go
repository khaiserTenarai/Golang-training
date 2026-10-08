package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ==========================================
// 1. DOMAIN MODELS & DTOs
// ==========================================

type contextKey string

const (
	RequestIDKey contextKey = "requestID"
	UserKey      contextKey = "user"
	RoleKey      contextKey = "role"
)

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateEmployeeDTO struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func (dto *CreateEmployeeDTO) Validate() error {
	if strings.TrimSpace(dto.Name) == "" {
		return errors.New("field 'name' is required")
	}
	if !strings.Contains(dto.Email, "@") {
		return errors.New("field 'email' must be a valid email address")
	}
	if strings.TrimSpace(dto.Department) == "" {
		return errors.New("field 'department' is required")
	}
	if dto.Salary <= 0 {
		return errors.New("field 'salary' must be greater than zero")
	}
	return nil
}

type QueryParams struct {
	Department string
	SortBy     string
	Order      string
	Page       int
	Limit      int
}

type PaginatedResponse struct {
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	TotalItems int        `json:"total_items"`
	TotalPages int        `json:"total_pages"`
	Data       []Employee `json:"data"`
}

// ==========================================
// 2. REPOSITORY LAYER (Data Access)
// ==========================================

type EmployeeRepository interface {
	Ping(ctx context.Context) error
	Create(ctx context.Context, emp *Employee) error
	GetByID(ctx context.Context, id int) (*Employee, error)
	List(ctx context.Context, params QueryParams) ([]Employee, int, error)
}

type inMemoryRepo struct {
	mu        sync.RWMutex
	employees map[int]Employee
	nextID    int
}

func NewInMemoryRepo() EmployeeRepository {
	repo := &inMemoryRepo{
		employees: make(map[int]Employee),
		nextID:    1,
	}
	// Seed Initial Data
	repo.employees[1] = Employee{ID: 1, Name: "Alice Smith", Email: "alice@company.com", Department: "Engineering", Salary: 95000, CreatedAt: time.Now()}
	repo.employees[2] = Employee{ID: 2, Name: "Bob Jones", Email: "bob@company.com", Department: "Design", Salary: 75000, CreatedAt: time.Now()}
	repo.employees[3] = Employee{ID: 3, Name: "Charlie Brown", Email: "charlie@company.com", Department: "Engineering", Salary: 105000, CreatedAt: time.Now()}
	repo.nextID = 4
	return repo
}

func (r *inMemoryRepo) Ping(ctx context.Context) error {
	return nil // Memory DB always ready
}

func (r *inMemoryRepo) Create(ctx context.Context, emp *Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	emp.ID = r.nextID
	emp.CreatedAt = time.Now()
	r.nextID++
	r.employees[emp.ID] = *emp
	return nil
}

func (r *inMemoryRepo) GetByID(ctx context.Context, id int) (*Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	emp, exists := r.employees[id]
	if !exists {
		return nil, errors.New("employee not found")
	}
	return &emp, nil
}

func (r *inMemoryRepo) List(ctx context.Context, params QueryParams) ([]Employee, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []Employee
	for _, emp := range r.employees {
		if params.Department == "" || strings.EqualFold(emp.Department, params.Department) {
			filtered = append(filtered, emp)
		}
	}

	// Sorting
	sort.Slice(filtered, func(i, j int) bool {
		switch params.SortBy {
		case "salary":
			if strings.EqualFold(params.Order, "desc") {
				return filtered[i].Salary > filtered[j].Salary
			}
			return filtered[i].Salary < filtered[j].Salary
		case "name":
			if strings.EqualFold(params.Order, "desc") {
				return filtered[i].Name > filtered[j].Name
			}
			return filtered[i].Name < filtered[j].Name
		default:
			return filtered[i].ID < filtered[j].ID
		}
	})

	totalItems := len(filtered)

	// Pagination
	startIndex := (params.Page - 1) * params.Limit
	if startIndex >= totalItems {
		return []Employee{}, totalItems, nil
	}

	endIndex := startIndex + params.Limit
	if endIndex > totalItems {
		endIndex = totalItems
	}

	return filtered[startIndex:endIndex], totalItems, nil
}

// ==========================================
// 3. SERVICE LAYER (Business Logic)
// ==========================================

type EmployeeService interface {
	CheckHealth(ctx context.Context) error
	CreateEmployee(ctx context.Context, dto CreateEmployeeDTO) (*Employee, error)
	GetEmployee(ctx context.Context, id int) (*Employee, error)
	ListEmployees(ctx context.Context, params QueryParams) (*PaginatedResponse, error)
}

type employeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CheckHealth(ctx context.Context) error {
	return s.repo.Ping(ctx)
}

func (s *employeeService) CreateEmployee(ctx context.Context, dto CreateEmployeeDTO) (*Employee, error) {
	if err := dto.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	emp := &Employee{
		Name:       dto.Name,
		Email:      dto.Email,
		Department: dto.Department,
		Salary:     dto.Salary,
	}

	if err := s.repo.Create(ctx, emp); err != nil {
		return nil, err
	}

	return emp, nil
}

func (s *employeeService) GetEmployee(ctx context.Context, id int) (*Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *employeeService) ListEmployees(ctx context.Context, params QueryParams) (*PaginatedResponse, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 || params.Limit > 100 {
		params.Limit = 10
	}

	items, totalItems, err := s.repo.List(ctx, params)
	if err != nil {
		return nil, err
	}

	totalPages := (totalItems + params.Limit - 1) / params.Limit

	return &PaginatedResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       items,
	}, nil
}

// ==========================================
// 4. HANDLER LAYER (HTTP Transport)
// ==========================================

type EmployeeHandler struct {
	svc EmployeeService
}

func NewEmployeeHandler(svc EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}

func (h *EmployeeHandler) Health(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"status": "UP", "timestamp": time.Now().Format(time.RFC3339)})
}

func (h *EmployeeHandler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.CheckHealth(r.Context()); err != nil {
		respondError(w, http.StatusServiceUnavailable, "Database or repository not ready")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"status": "READY"})
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateEmployeeDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	emp, err := h.svc.CreateEmployee(r.Context(), dto)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, emp)
}

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	emp, err := h.svc.GetEmployee(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, emp)
}

func (h *EmployeeHandler) List(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	params := QueryParams{
		Department: r.URL.Query().Get("department"),
		SortBy:     r.URL.Query().Get("sort_by"),
		Order:      r.URL.Query().Get("order"),
		Page:       page,
		Limit:      limit,
	}

	resp, err := h.svc.ListEmployees(r.Context(), params)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to list employees")
		return
	}

	respondJSON(w, http.StatusOK, resp)
}

// Helpers
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// ==========================================
// 5. MIDDLEWARE PIPELINE
// ==========================================

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			b := make([]byte, 8)
			rand.Read(b)
			reqID = fmt.Sprintf("%x", b)
		}
		w.Header().Set("X-Request-ID", reqID)
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sw, r)

		reqID, _ := r.Context().Value(RequestIDKey).(string)
		log.Printf("[REQ-ID: %s] %s %s %d - %v", reqID, r.Method, r.URL.Path, sw.status, time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			respondError(w, http.StatusUnauthorized, "Missing or invalid authorization header")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		
		// Simulated JWT parsing logic:
		// "admin-token" -> Admin user
		// "user-token"  -> Regular user
		var user, role string
		switch token {
		case "admin-token":
			user, role = "admin_user", "Admin"
		case "user-token":
			user, role = "staff_user", "Staff"
		default:
			respondError(w, http.StatusUnauthorized, "Invalid bearer token")
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user)
		ctx = context.WithValue(ctx, RoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireRole(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(RoleKey).(string)
		if !ok || role != requiredRole {
			respondError(w, http.StatusForbidden, "Forbidden: insufficient permissions")
			return
		}
		next(w, r)
	}
}

// ==========================================
// 6. MAIN APPLICATION BOOTSTRAP (DI)
// ==========================================

func main() {
	// 1. Instantiate Infrastructure/Repository Layer
	repo := NewInMemoryRepo()

	// 2. Instantiate Service Layer (Inject Repo)
	svc := NewEmployeeService(repo)

	// 3. Instantiate Handler Layer (Inject Service)
	handler := NewEmployeeHandler(svc)

	// 4. Setup Router
	mux := http.NewServeMux()

	// Unauthenticated Health Endpoints
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)

	// Authenticated Protected Routes
	mux.Handle("GET /employees", AuthMiddleware(http.HandlerFunc(handler.List)))
	mux.Handle("GET /employees/{id}", AuthMiddleware(http.HandlerFunc(handler.GetByID)))
	
	// Authenticated + Role-Restricted Admin Route
	mux.Handle("POST /employees", AuthMiddleware(RequireRole("Admin", handler.Create)))

	// Global Middleware Wrapping
	siteHandler := RequestIDMiddleware(LoggerMiddleware(mux))

	fmt.Println("Production Layered Server listening on http://localhost:8080")
	fmt.Println("Try visiting /health or calling GET /employees with header 'Authorization: Bearer user-token'")
	log.Fatal(http.ListenAndServe(":8080", siteHandler))
}