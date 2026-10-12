// Day 12 - Task 2: Separate Handler and Service
//
// The handler only deals with HTTP (reading the JSON, picking the
// status code). The service owns the actual rules - here, a minimum
// salary that you choose at startup, and no duplicate emails - and has
// no idea that HTTP even exists.
//   curl -X POST http://localhost:8080/employee \
//     -d '{"name":"Amit","email":"amit@example.com","salary":40000}'
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

var (
	errDuplicateEmail = errors.New("an employee with that email already exists")
	errSalaryTooLow   = errors.New("salary is below the minimum")
)

// ---------- service: business rules, no HTTP in sight ----------

type EmployeeService struct {
	mu        sync.Mutex
	employees []Employee
	nextID    int
	minSalary float64
}

func NewEmployeeService(minSalary float64) *EmployeeService {
	return &EmployeeService{nextID: 1, minSalary: minSalary}
}

func (s *EmployeeService) Add(e Employee) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if e.Salary < s.minSalary {
		return Employee{}, fmt.Errorf("%w (minimum is %.2f)", errSalaryTooLow, s.minSalary)
	}
	for _, existing := range s.employees {
		if strings.EqualFold(existing.Email, e.Email) {
			return Employee{}, errDuplicateEmail
		}
	}

	e.ID = s.nextID
	s.nextID++
	s.employees = append(s.employees, e)
	return e, nil
}

func (s *EmployeeService) All() []Employee {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]Employee, len(s.employees))
	copy(result, s.employees)
	return result
}

// ---------- handler: only HTTP ----------

type EmployeeHandler struct {
	service *EmployeeService
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *EmployeeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.All())

	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}

		created, err := h.service.Add(e)
		if errors.Is(err, errDuplicateEmail) {
			writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
			return
		}
		if errors.Is(err, errSalaryTooLow) {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusCreated, created)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Minimum salary the service should accept: ")
	salaryText, _ := reader.ReadString('\n')
	minSalary, err := strconv.ParseFloat(strings.TrimSpace(salaryText), 64)
	if err != nil || minSalary < 0 {
		fmt.Println("That wasn't a valid amount, using 0.")
		minSalary = 0
	}

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	handler := &EmployeeHandler{service: NewEmployeeService(minSalary)}
	http.HandleFunc("/employee", handler.Handle)

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
