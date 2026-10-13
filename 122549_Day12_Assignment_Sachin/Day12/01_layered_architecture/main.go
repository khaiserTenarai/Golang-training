// Day 12 - Task 1: Create layered architecture
//
// The request travels down through three layers and the answer comes
// back up the same way:
//
//   handler  (HTTP in/out)  ->  service (business rules)  ->  repository (data)
//
// Each layer only talks to the one directly below it.
//   GET  /employee, GET /employee/{id}, POST /employee
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"layered/handler"
	"layered/repository"
	"layered/service"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Port to run the server on (press Enter for 8080): ")
	text, _ := reader.ReadString('\n')
	port := strings.TrimSpace(text)
	if port == "" {
		port = "8080"
	}

	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo)
	h := handler.NewEmployeeHandler(svc)

	http.HandleFunc("/employee", h.Employees)
	http.HandleFunc("/employee/", h.EmployeeByID)

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
