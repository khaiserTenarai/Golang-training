package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	Create(ctx context.Context, e *Employee) error
	GetByID(ctx context.Context, id int) (*Employee, error)
	GetAll(ctx context.Context) ([]*Employee, error)
	Update(ctx context.Context, e *Employee) error
	Delete(ctx context.Context, id int) error
}

type postgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &postgresEmployeeRepository{db: db}
}

func (r *postgresEmployeeRepository) Create(ctx context.Context, e *Employee) error {
	return r.db.QueryRow(ctx, `INSERT INTO employees (name, email, salary) VALUES ($1, $2, $3) RETURNING id`,
		e.Name, e.Email, e.Salary).Scan(&e.ID)
}

func (r *postgresEmployeeRepository) GetByID(ctx context.Context, id int) (*Employee, error) {
	e := &Employee{}
	err := r.db.QueryRow(ctx, `SELECT id, name, email, salary FROM employees WHERE id = $1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *postgresEmployeeRepository) GetAll(ctx context.Context) ([]*Employee, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email, salary FROM employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*Employee
	for rows.Next() {
		e := &Employee{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, rows.Err()
}

func (r *postgresEmployeeRepository) Update(ctx context.Context, e *Employee) error {
	result, err := r.db.Exec(ctx, `UPDATE employees SET name = $1, email = $2, salary = $3 WHERE id = $4`,
		e.Name, e.Email, e.Salary, e.ID)
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

type EmployeeService interface {
	CreateEmployee(ctx context.Context, e *Employee) error
	GetEmployee(ctx context.Context, id int) (*Employee, error)
	ListEmployees(ctx context.Context) ([]*Employee, error)
	UpdateEmployee(ctx context.Context, e *Employee) error
	DeleteEmployee(ctx context.Context, id int) error
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

func (s *employeeService) ListEmployees(ctx context.Context) ([]*Employee, error) {
	return s.repo.GetAll(ctx)
}

func (s *employeeService) UpdateEmployee(ctx context.Context, e *Employee) error {
	return s.repo.Update(ctx, e)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
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
		employees, err := c.service.ListEmployees(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(employees)
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
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *EmployeeController) HandleItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idStr)
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
		json.NewEncoder(w).Encode(e)
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
		json.NewEncoder(w).Encode(e)
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

func main() {
	connString := "postgres://postgres:password@localhost:5432/rsadb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewPostgresEmployeeRepository(db)
	service := NewEmployeeService(repo)
	controller := NewEmployeeController(service)

	http.HandleFunc("/employees", controller.HandleCollection)
	http.HandleFunc("/employees/", controller.HandleItem)

	http.ListenAndServe(":8080", nil)
}
