// Day 14 - Task 13: Serve static files
//
// router.Static("/static", folder) serves everything in a folder
// (HTML, CSS, images, JS) under a URL prefix. The home page is the
// static index.html, and that page loads the employee list from the
// JSON API underneath it.
//
// Run it from inside this folder:   go run .
// (the first run needs "go mod tidy" once, in the Day14 folder above)
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	folder := ask(reader, "Which folder holds the static files? (press Enter for ./static): ")
	if folder == "" {
		folder = "static"
	}
	if _, err := os.Stat(filepath.Join(folder, "index.html")); err != nil {
		fmt.Println("Couldn't find index.html inside", folder, "- run this from the 13_static_files folder.")
		return
	}

	employees := askEmployees(reader)
	port := askPort(reader)

	router := gin.Default()

	router.Static("/static", folder)

	router.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(folder, "index.html"))
	})

	router.GET("/api/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	fmt.Println("Open http://localhost:" + port + " in your browser")
	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
