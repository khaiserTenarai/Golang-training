package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"sync"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var (
	employees          = []Employee{}
	nextID             = 1
	mu                 sync.Mutex
	allowedDepartments []string
)

// ---------- the validation layer ----------

func validateEmployee(e Employee) map[string]string {
	problems := map[string]string{}

	name := strings.TrimSpace(e.Name)
	if len(name) < 2 || len(name) > 50 {
		problems["name"] = "must be between 2 and 50 characters"
	}

	if addr, err := mail.ParseAddress(e.Email); err != nil || addr.Address != e.Email {
		problems["email"] = "is not a valid email address"
	}

	validDepartment := false
	for _, d := range allowedDepartments {
		if strings.EqualFold(d, e.Department) {
			validDepartment = true
		}
	}
	if !validDepartment {
		problems["department"] = "must be one of: " + strings.Join(allowedDepartments, ", ")
	}

	if e.Salary <= 0 || e.Salary > 10000000 {
		problems["salary"] = "must be above 0 and at most 10,000,000"
	}

	return problems
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		writeJSON(w, http.StatusOK, employees)

	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}

		if problems := validateEmployee(e); len(problems) > 0 {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]interface{}{"errors": problems})
			return
		}

		mu.Lock()
		defer mu.Unlock()
		e.ID = nextID
		nextID++
		employees = append(employees, e)
		writeJSON(w, http.StatusCreated, e)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Allowed departments, separated by commas (for example IT,HR,Finance): ")
	line, _ := reader.ReadString('\n')
	for _, d := range strings.Split(line, ",") {
		d = strings.TrimSpace(d)
		if d != "" {
			allowedDepartments = append(allowedDepartments, d)
		}
	}
	if len(allowedDepartments) == 0 {
		fmt.Println("Nothing entered, using IT and HR.")
		allowedDepartments = []string{"IT", "HR"}
	}

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/employee", employeesHandler)

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
