package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Employee struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Salary float64
}

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, e *Employee) error {
	query := `INSERT INTO employees (name, email, age, salary) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(ctx, query, e.Name, e.Email, e.Age, e.Salary).Scan(&e.ID)
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id int) (*Employee, error) {
	query := `SELECT id, name, email, age, salary FROM employees WHERE id = $1`
	e := &Employee{}
	err := r.db.QueryRow(ctx, query, id).Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *EmployeeRepository) GetAll(ctx context.Context) ([]*Employee, error) {
	query := `SELECT id, name, email, age, salary FROM employees ORDER BY id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*Employee
	for rows.Next() {
		e := &Employee{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, rows.Err()
}

func (r *EmployeeRepository) Update(ctx context.Context, e *Employee) error {
	query := `UPDATE employees SET name = $1, email = $2, age = $3, salary = $4 WHERE id = $5`
	result, err := r.db.Exec(ctx, query, e.Name, e.Email, e.Age, e.Salary, e.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *EmployeeRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM employees WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func connect(connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return pool, nil
}

func main() {
	connString := "postgres://postgres:password@localhost:5432/employeedb"

	db, err := connect(connString)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	repo := NewEmployeeRepository(db)
	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Employee")
		fmt.Println("2. Get Employee by ID")
		fmt.Println("3. List All Employees")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		choice := readLine(reader)

		switch choice {
		case "1":
			createEmployee(ctx, repo, reader)
		case "2":
			getEmployee(ctx, repo, reader)
		case "3":
			listEmployees(ctx, repo)
		case "4":
			updateEmployee(ctx, repo, reader)
		case "5":
			deleteEmployee(ctx, repo, reader)
		case "6":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func createEmployee(ctx context.Context, repo *EmployeeRepository, reader *bufio.Reader) {
	fmt.Print("Name: ")
	name := readLine(reader)

	fmt.Print("Email: ")
	email := readLine(reader)

	fmt.Print("Age: ")
	age, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Salary: ")
	salary, _ := strconv.ParseFloat(readLine(reader), 64)

	e := &Employee{Name: name, Email: email, Age: age, Salary: salary}
	if err := repo.Create(ctx, e); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created employee with ID:", e.ID)
}

func getEmployee(ctx context.Context, repo *EmployeeRepository, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	e, err := repo.GetByID(ctx, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	printEmployee(e)
}

func listEmployees(ctx context.Context, repo *EmployeeRepository) {
	employees, err := repo.GetAll(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}
	for _, e := range employees {
		printEmployee(e)
	}
}

func updateEmployee(ctx context.Context, repo *EmployeeRepository, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	existing, err := repo.GetByID(ctx, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name [%s]: ", existing.Name)
	if name := readLine(reader); name != "" {
		existing.Name = name
	}

	fmt.Printf("Email [%s]: ", existing.Email)
	if email := readLine(reader); email != "" {
		existing.Email = email
	}

	fmt.Printf("Age [%d]: ", existing.Age)
	if ageStr := readLine(reader); ageStr != "" {
		existing.Age, _ = strconv.Atoi(ageStr)
	}

	fmt.Printf("Salary [%.2f]: ", existing.Salary)
	if salaryStr := readLine(reader); salaryStr != "" {
		existing.Salary, _ = strconv.ParseFloat(salaryStr, 64)
	}

	if err := repo.Update(ctx, existing); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Employee updated")
}

func deleteEmployee(ctx context.Context, repo *EmployeeRepository, reader *bufio.Reader) {
	fmt.Print("Employee ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	if err := repo.Delete(ctx, id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Employee deleted")
}

func printEmployee(e *Employee) {
	fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d | Salary: %.2f\n",
		e.ID, e.Name, e.Email, e.Age, e.Salary)
}
