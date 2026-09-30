package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type PaginatedEmployees struct {
	Data     []*Employee `json:"data"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

var ErrNotFound = errors.New("employee not found")
var ErrInvalidCredentials = errors.New("invalid credentials")

type EmployeeRepository interface {
	Create(ctx context.Context, e *Employee) error
	GetByID(ctx context.Context, id int) (*Employee, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*Employee, error)
	Update(ctx context.Context, e *Employee) error
	Delete(ctx context.Context, id int) error
	UpdateSalary(ctx context.Context, id int, newSalary float64) error
}

type postgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &postgresEmployeeRepository{db: db}
}

func (r *postgresEmployeeRepository) Create(ctx context.Context, e *Employee) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`,
		e.Name, e.Email, e.Department, e.Salary).Scan(&e.ID)
}

func (r *postgresEmployeeRepository) GetByID(ctx context.Context, id int) (*Employee, error) {
	e := &Employee{}
	err := r.db.QueryRow(ctx,
		`SELECT id, name, email, department, salary FROM employees WHERE id = $1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *postgresEmployeeRepository) GetAll(ctx context.Context, page, pageSize int) ([]*Employee, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, email, department, salary FROM employees ORDER BY id LIMIT $1 OFFSET $2`,
		pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*Employee
	for rows.Next() {
		e := &Employee{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, rows.Err()
}

func (r *postgresEmployeeRepository) Update(ctx context.Context, e *Employee) error {
	result, err := r.db.Exec(ctx,
		`UPDATE employees SET name = $1, email = $2, department = $3 WHERE id = $4`,
		e.Name, e.Email, e.Department, e.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *postgresEmployeeRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *postgresEmployeeRepository) UpdateSalary(ctx context.Context, id int, newSalary float64) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var oldSalary float64
	err = tx.QueryRow(ctx, `SELECT salary FROM employees WHERE id = $1 FOR UPDATE`, id).Scan(&oldSalary)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, `UPDATE employees SET salary = $1 WHERE id = $2`, newSalary, id); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx,
		`INSERT INTO salary_history (employee_id, old_salary, new_salary) VALUES ($1, $2, $3)`,
		id, oldSalary, newSalary); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

type EmployeeService interface {
	CreateEmployee(ctx context.Context, e *Employee) error
	GetEmployee(ctx context.Context, id int) (*Employee, error)
	ListEmployees(ctx context.Context, page, pageSize int) ([]*Employee, error)
	UpdateEmployee(ctx context.Context, e *Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	UpdateSalary(ctx context.Context, id int, newSalary float64) error
}

type employeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(ctx context.Context, e *Employee) error {
	return s.repo.Create(ctx, e)
}

func (s *employeeService) GetEmployee(ctx context.Context, id int) (*Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *employeeService) ListEmployees(ctx context.Context, page, pageSize int) ([]*Employee, error) {
	page, pageSize = normalizePagination(page, pageSize)
	return s.repo.GetAll(ctx, page, pageSize)
}

func (s *employeeService) UpdateEmployee(ctx context.Context, e *Employee) error {
	return s.repo.Update(ctx, e)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *employeeService) UpdateSalary(ctx context.Context, id int, newSalary float64) error {
	return s.repo.UpdateSalary(ctx, id, newSalary)
}

func normalizePagination(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return page, pageSize
}

type AuthService struct {
	db *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) *AuthService {
	return &AuthService{db: db}
}

func (a *AuthService) Register(ctx context.Context, username, password, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = a.db.Exec(ctx,
		`INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`,
		username, string(hash), role)
	return err
}

func (a *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	var storedHash, role string
	err := a.db.QueryRow(ctx,
		`SELECT password_hash, role FROM users WHERE username = $1`, username).
		Scan(&storedHash, &role)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}
	return role, nil
}

func (a *AuthService) UserExists(ctx context.Context, username string) bool {
	var exists bool
	a.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`, username).Scan(&exists)
	return exists
}

type EmployeeController struct {
	service EmployeeService
}

func NewEmployeeController(service EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) HandleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
		employees, err := c.service.ListEmployees(r.Context(), page, pageSize)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		page, pageSize = normalizePagination(page, pageSize)
		writeJSON(w, http.StatusOK, PaginatedEmployees{Data: employees, Page: page, PageSize: pageSize})
	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := c.service.CreateEmployee(r.Context(), &e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusCreated, e)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *EmployeeController) HandleItem(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/employees/")

	if strings.HasSuffix(path, "/salary") {
		idStr := strings.TrimSuffix(path, "/salary")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c.handleSalaryUpdate(w, r, id)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		e, err := c.service.GetEmployee(r.Context(), id)
		if errors.Is(err, ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, e)
	case http.MethodPut:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		e.ID = id
		if err := c.service.UpdateEmployee(r.Context(), &e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, e)
	case http.MethodDelete:
		if err := c.service.DeleteEmployee(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *EmployeeController) handleSalaryUpdate(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		Salary float64 `json:"salary"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.UpdateSalary(r.Context(), id, payload.Salary); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type AuthController struct {
	auth *AuthService
}

func NewAuthController(auth *AuthService) *AuthController {
	return &AuthController{auth: auth}
}

func (c *AuthController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if payload.Role == "" {
		payload.Role = "user"
	}
	if err := c.auth.Register(r.Context(), payload.Username, payload.Password, payload.Role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	role, err := c.auth.Login(r.Context(), payload.Username, payload.Password)
	if errors.Is(err, ErrInvalidCredentials) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"username": payload.Username, "role": role})
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(auth *AuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("X-Username")
		if username == "" || !auth.UserExists(r.Context(), username) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	connString := "postgres://postgres:password@localhost:5432/capstonedb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewPostgresEmployeeRepository(db)
	service := NewEmployeeService(repo)
	employeeController := NewEmployeeController(service)

	authService := NewAuthService(db)
	authController := NewAuthController(authService)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", authController.HandleRegister)
	mux.HandleFunc("/login", authController.HandleLogin)
	mux.Handle("/employees", authMiddleware(authService, http.HandlerFunc(employeeController.HandleCollection)))
	mux.Handle("/employees/", authMiddleware(authService, http.HandlerFunc(employeeController.HandleItem)))

	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
