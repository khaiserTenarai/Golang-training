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
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit < 1 {
			limit = 5
		}
		if limit > 50 {
			limit = 50
		}

		total := len(employees)
		start := (page - 1) * limit
		if start > total {
			start = total
		}
		end := start + limit
		if end > total {
			end = total
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"data":        employees[start:end],
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + limit - 1) / limit,
		})
	})

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
