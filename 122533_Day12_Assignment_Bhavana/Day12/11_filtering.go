// Day 12 - Task 11: Implement filtering
//
// Narrow the list down with query parameters - they can be combined:
//   GET /employee?department=IT
//   GET /employee?min_salary=35000&max_salary=50000
//   GET /employee?search=employee 1
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func askPort(reader *bufio.Reader) string {
	port := ask(reader, "Port to run the server on (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}
	return port
}

func askCount(reader *bufio.Reader) int {
	count, err := strconv.Atoi(ask(reader, "How many sample employees should be generated? "))
	if err != nil || count < 0 {
		fmt.Println("That's not a valid number, generating 10.")
		count = 10
	}
	return count
}

// sampleEmployees makes up `count` employees spread over a few departments.
func sampleEmployees(count int) []Employee {
	departments := []string{"IT", "HR", "Finance", "Sales"}
	list := make([]Employee, 0, count)
	for i := 1; i <= count; i++ {
		list = append(list, Employee{
			ID:         i,
			Name:       fmt.Sprintf("Employee %d", i),
			Email:      fmt.Sprintf("employee%d@example.com", i),
			Department: departments[i%len(departments)],
			Salary:     float64(30000 + (i*7919)%40000),
		})
	}
	return list
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	employees := sampleEmployees(askCount(reader))
	port := askPort(reader)

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		department := query.Get("department")
		search := strings.ToLower(query.Get("search"))
		minSalary, _ := strconv.ParseFloat(query.Get("min_salary"), 64)
		maxSalary, _ := strconv.ParseFloat(query.Get("max_salary"), 64)

		results := []Employee{}
		for _, e := range employees {
			if department != "" && !strings.EqualFold(e.Department, department) {
				continue
			}
			if search != "" && !strings.Contains(strings.ToLower(e.Name), search) {
				continue
			}
			if minSalary > 0 && e.Salary < minSalary {
				continue
			}
			if maxSalary > 0 && e.Salary > maxSalary {
				continue
			}
			results = append(results, e)
		}
		writeJSON(w, http.StatusOK, results)
	})

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
