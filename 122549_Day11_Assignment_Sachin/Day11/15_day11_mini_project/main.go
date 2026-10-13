package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var db *sql.DB

func ask(reader *bufio.Reader, label, fallback string) string {
	fmt.Printf("%s [%s]: ", label, fallback)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return fallback
	}
	return text
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func validate(e Employee) []string {
	var problems []string
	if strings.TrimSpace(e.Name) == "" {
		problems = append(problems, "name is required")
	}
	if !strings.Contains(e.Email, "@") || !strings.Contains(e.Email, ".") {
		problems = append(problems, "email must look like name@example.com")
	}
	if strings.TrimSpace(e.Department) == "" {
		problems = append(problems, "department is required")
	}
	if e.Salary <= 0 {
		problems = append(problems, "salary must be greater than 0")
	}
	return problems
}

// 23505 is Postgres' error code for "unique constraint violated"
func isDuplicateEmail(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == "23505"
}

// ---------- middleware ----------

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)
		log.Printf("%s %s -> %d (%v)", r.Method, r.URL.Path, recorder.status, time.Since(start))
	})
}

// ---------- handlers ----------

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listEmployees(w, r)
	case http.MethodPost:
		createEmployee(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getEmployee(w, id)
	case http.MethodPut:
		updateEmployee(w, r, id)
	case http.MethodDelete:
		deleteEmployee(w, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func listEmployees(w http.ResponseWriter, r *http.Request) {
	department := r.URL.Query().Get("department")
	name := r.URL.Query().Get("name")

	query := "SELECT id, name, email, department, salary FROM employees WHERE 1=1"
	var args []interface{}

	if department != "" {
		args = append(args, department)
		query += fmt.Sprintf(" AND department ILIKE $%d", len(args))
	}
	if name != "" {
		args = append(args, "%"+name+"%")
		query += fmt.Sprintf(" AND name ILIKE $%d", len(args))
	}
	query += " ORDER BY id"

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("list failed:", err)
		writeError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			log.Println("scan failed:", err)
			writeError(w, http.StatusInternalServerError, "something went wrong")
			return
		}
		employees = append(employees, e)
	}
	writeJSON(w, http.StatusOK, employees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		writeError(w, http.StatusBadRequest, "could not read the JSON body")
		return
	}
	if problems := validate(e); len(problems) > 0 {
		writeJSON(w, http.StatusBadRequest, map[string][]string{"errors": problems})
		return
	}

	err := db.QueryRow(
		`INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`,
		e.Name, e.Email, e.Department, e.Salary,
	).Scan(&e.ID)
	if isDuplicateEmail(err) {
		writeError(w, http.StatusConflict, "that email is already in use")
		return
	}
	if err != nil {
		log.Println("insert failed:", err)
		writeError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	writeJSON(w, http.StatusCreated, e)
}

func getEmployee(w http.ResponseWriter, id int) {
	var e Employee
	err := db.QueryRow(`SELECT id, name, email, department, salary FROM employees WHERE id = $1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	if err != nil {
		log.Println("get failed:", err)
		writeError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	writeJSON(w, http.StatusOK, e)
}

func updateEmployee(w http.ResponseWriter, r *http.Request, id int) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		writeError(w, http.StatusBadRequest, "could not read the JSON body")
		return
	}
	if problems := validate(e); len(problems) > 0 {
		writeJSON(w, http.StatusBadRequest, map[string][]string{"errors": problems})
		return
	}

	result, err := db.Exec(
		`UPDATE employees SET name = $1, email = $2, department = $3, salary = $4 WHERE id = $5`,
		e.Name, e.Email, e.Department, e.Salary, id,
	)
	if isDuplicateEmail(err) {
		writeError(w, http.StatusConflict, "that email is already in use")
		return
	}
	if err != nil {
		log.Println("update failed:", err)
		writeError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	e.ID = id
	writeJSON(w, http.StatusOK, e)
}

func deleteEmployee(w http.ResponseWriter, id int) {
	result, err := db.Exec(`DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		log.Println("delete failed:", err)
		writeError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	host := ask(reader, "Database host", "localhost")
	dbPort := ask(reader, "Database port", "5432")
	user := ask(reader, "Database user", "postgres")
	password := ask(reader, "Database password", "postgres")
	dbName := ask(reader, "Database name", "employee_db")
	serverPort := ask(reader, "Port for this server", "8080")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbName)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not set up the connection:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Could not reach PostgreSQL:", err)
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		department VARCHAR(100) NOT NULL,
		salary NUMERIC(12,2) NOT NULL
	)`)
	if err != nil {
		fmt.Println("Could not create the table:", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employeesHandler)
	mux.HandleFunc("/employee/", employeeHandler)

	fmt.Println("Connected. Employee API running on port", serverPort)
	if err := http.ListenAndServe(":"+serverPort, loggingMiddleware(mux)); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
