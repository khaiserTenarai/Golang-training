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

func askEmployees(reader *bufio.Reader) []Employee {
	count, err := strconv.Atoi(ask(reader, "How many employees do you want to start with? "))
	if err != nil || count < 0 {
		fmt.Println("That's not a valid number, starting with no employees.")
		count = 0
	}

	list := []Employee{}
	for i := 1; i <= count; i++ {
		fmt.Printf("\n--- Employee %d ---\n", i)
		name := ask(reader, "Name: ")
		email := ask(reader, "Email: ")
		department := ask(reader, "Department: ")
		salary, _ := strconv.ParseFloat(ask(reader, "Salary: "), 64)
		list = append(list, Employee{ID: i, Name: name, Email: email, Department: department, Salary: salary})
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
	employees := askEmployees(reader)
	port := askPort(reader)

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, employees)
	})

	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		idText := strings.TrimPrefix(r.URL.Path, "/employee/")
		id, err := strconv.Atoi(idText)
		if err != nil {
			http.Error(w, "the id in the URL has to be a number", http.StatusBadRequest)
			return
		}

		for _, e := range employees {
			if e.ID == id {
				writeJSON(w, http.StatusOK, e)
				return
			}
		}
		http.Error(w, "no employee with that id", http.StatusNotFound)
	})

	fmt.Println("Try it: http://localhost:" + port + "/employee/1")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
