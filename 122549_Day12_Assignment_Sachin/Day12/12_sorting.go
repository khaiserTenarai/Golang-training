package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
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
		sortBy := r.URL.Query().Get("sort_by")
		if sortBy == "" {
			sortBy = "id"
		}
		descending := strings.EqualFold(r.URL.Query().Get("order"), "desc")

		var less func(a, b Employee) bool
		switch sortBy {
		case "id":
			less = func(a, b Employee) bool { return a.ID < b.ID }
		case "name":
			less = func(a, b Employee) bool { return a.Name < b.Name }
		case "department":
			less = func(a, b Employee) bool { return a.Department < b.Department }
		case "salary":
			less = func(a, b Employee) bool { return a.Salary < b.Salary }
		default:
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "sort_by must be id, name, department or salary"})
			return
		}

		// sort a copy, so the original order isn't disturbed for the next request
		sorted := make([]Employee, len(employees))
		copy(sorted, employees)
		sort.SliceStable(sorted, func(i, j int) bool {
			if descending {
				return less(sorted[j], sorted[i])
			}
			return less(sorted[i], sorted[j])
		})

		writeJSON(w, http.StatusOK, sorted)
	})

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
