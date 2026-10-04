package main

import (
	"context"
	"database/sql"
	"encoding/json"
	
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// ==========================================
// MODELS
// ==========================================

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

// ==========================================
// REPOSITORY LAYER
// ==========================================

type AppRepository interface {
	RegisterUser(username, hash string) error
	GetUserByUsername(username string) (*User, error)
	CreateEmployee(emp *Employee) error
	GetEmployees(limit, offset int) ([]Employee, error)
	UpdateSalaryTx(ctx context.Context, empID int, newSalary float64) error
}

type repo struct {
	db *sql.DB
}

func (r *repo) RegisterUser(username, hash string) error {
	_, err := r.db.Exec(`INSERT INTO users (username, password_hash) VALUES ($1, $2)`, username, hash)
	return err
}

func (r *repo) GetUserByUsername(username string) (*User, error) {
	u := &User{}
	err := r.db.QueryRow(`SELECT id, username, password_hash FROM users WHERE username = $1`, username).
		Scan(&u.ID, &u.Username, &u.PasswordHash)
	return u, err
}

func (r *repo) CreateEmployee(emp *Employee) error {
	return r.db.QueryRow(`INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3) RETURNING id`,
		emp.Name, emp.Position, emp.Salary).Scan(&emp.ID)
}

func (r *repo) GetEmployees(limit, offset int) ([]Employee, error) {
	rows, err := r.db.Query(`SELECT id, name, position, salary FROM employees ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []Employee
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Position, &e.Salary); err != nil {
			return nil, err
		}
		emps = append(emps, e)
	}
	return emps, nil
}

func (r *repo) UpdateSalaryTx(ctx context.Context, empID int, newSalary float64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var oldSalary float64
	err = tx.QueryRowContext(ctx, `SELECT salary FROM employees WHERE id = $1 FOR UPDATE`, empID).Scan(&oldSalary)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO salary_history (employee_id, old_salary, new_salary) VALUES ($1, $2, $3)`, empID, oldSalary, newSalary)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE employees SET salary = $1 WHERE id = $2`, newSalary, empID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// ==========================================
// SERVICE LAYER
// ==========================================

type AppService struct {
	repo AppRepository
}

func (s *AppService) Register(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.repo.RegisterUser(username, string(hash))
}

func (s *AppService) Authenticate(username, password string) error {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}

// ==========================================
// MIDDLEWARE (Logging & Auth)
// ==========================================

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func AuthMiddleware(svc *AppService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || svc.Authenticate(user, pass) != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// ==========================================
// CONTROLLER (Handlers)
// ==========================================

type Controller struct {
	svc  *AppService
	repo AppRepository
}

func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&input)
	
	if err := c.svc.Register(input.Username, input.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "user created"}`))
}

func (c *Controller) HandleEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var emp Employee
		json.NewDecoder(r.Body).Decode(&emp)
		c.repo.CreateEmployee(&emp)
		json.NewEncoder(w).Encode(emp)
		return
	}

	if r.Method == http.MethodGet {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 1 { page = 1 }
		limit := 5
		offset := (page - 1) * limit
		
		emps, _ := c.repo.GetEmployees(limit, offset)
		json.NewEncoder(w).Encode(emps)
		return
	}
}

func (c *Controller) UpdateSalary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var input struct {
		ID        int     `json:"id"`
		NewSalary float64 `json:"new_salary"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	err := c.repo.UpdateSalaryTx(r.Context(), input.ID, input.NewSalary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message": "salary updated and history recorded"}`))
}

// ==========================================
// MAIN APP WIRE-UP
// ==========================================

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=yourpassword dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := &repo{db: db}
	service := &AppService{repo: repository}
	controller := &Controller{svc: service, repo: repository}

	mux := http.NewServeMux()
	
	// Public Route
	mux.HandleFunc("/register", controller.RegisterUser)
	
	// Protected Routes (Require Basic Auth)
	mux.HandleFunc("/employees", AuthMiddleware(service, controller.HandleEmployees))
	mux.HandleFunc("/employees/salary", AuthMiddleware(service, controller.UpdateSalary))

	// Apply Logging Middleware globally
	handler := LoggingMiddleware(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}