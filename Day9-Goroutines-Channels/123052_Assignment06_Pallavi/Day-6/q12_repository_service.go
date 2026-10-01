package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:Info%402k26@localhost:5432/assignment_db"

// --- Model ---
type Employee struct {
	ID    int
	Name  string
	Email string
}

// --- 1. Repository Layer Interface & Implementation ---
type EmployeeRepository interface {
	Create(ctx context.Context, emp *Employee) error
	GetByID(ctx context.Context, id int) (*Employee, error)
}

type pgxEmployeeRepo struct {
	conn *pgx.Conn
}

func NewEmployeeRepo(conn *pgx.Conn) EmployeeRepository {
	return &pgxEmployeeRepo{conn: conn}
}

func (r *pgxEmployeeRepo) Create(ctx context.Context, emp *Employee) error {
	return r.conn.QueryRow(ctx, "INSERT INTO repo_employees (name, email) VALUES ($1, $2) RETURNING id", emp.Name, emp.Email).Scan(&emp.ID)
}

func (r *pgxEmployeeRepo) GetByID(ctx context.Context, id int) (*Employee, error) {
	var emp Employee
	err := r.conn.QueryRow(ctx, "SELECT id, name, email FROM repo_employees WHERE id=$1", id).Scan(&emp.ID, &emp.Name, &emp.Email)
	return &emp, err
}

// --- 2. Service Layer Interface & Implementation ---
type EmployeeService interface {
	RegisterEmployee(ctx context.Context, name, email string) (*Employee, error)
	FindEmployee(ctx context.Context, id int) (*Employee, error)
}

type employeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) RegisterEmployee(ctx context.Context, name, email string) (*Employee, error) {
	emp := &Employee{Name: name, Email: email}
	err := s.repo.Create(ctx, emp)
	return emp, err
}

func (s *employeeService) FindEmployee(ctx context.Context, id int) (*Employee, error) {
	return s.repo.GetByID(ctx, id)
}

// --- 3. Controller/Main Application ---
func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	_, _ = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS repo_employees (id SERIAL PRIMARY KEY, name VARCHAR(100), email VARCHAR(100) UNIQUE);`)

	// Dependency Injection Setup
	repo := NewEmployeeRepo(conn)
	service := NewEmployeeService(repo)

	// Invoke Business Logic via Controller / Service
	emp, err := service.RegisterEmployee(ctx, "Sarah Jenkins", "sarah@example.com")
	if err != nil {
		log.Fatalf("Service creation error: %v\n", err)
	}
	fmt.Printf("[Service Layer] Employee Registered: ID=%d, Name=%s\n", emp.ID, emp.Name)

	fetched, err := service.FindEmployee(ctx, emp.ID)
	if err != nil {
		log.Fatalf("Service fetch error: %v\n", err)
	}
	fmt.Printf("[Service Layer] Employee Fetched: %s (%s)\n", fetched.Name, fetched.Email)
}