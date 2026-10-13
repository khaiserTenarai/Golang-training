package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"employeeapi/handler"
	"employeeapi/middleware"
	"employeeapi/repository"
	"employeeapi/service"
)

func ask(reader *bufio.Reader, label, fallback string) string {
	fmt.Printf("%s [%s]: ", label, fallback)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return fallback
	}
	return text
}

func askRequired(reader *bufio.Reader, label string) string {
	for {
		fmt.Printf("%s: ", label)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Database (press Enter to accept the value in brackets) ---")
	host := ask(reader, "Host", "localhost")
	dbPort := ask(reader, "Port", "5432")
	user := ask(reader, "User", "postgres")
	password := ask(reader, "Password", "postgres")
	dbName := ask(reader, "Database name", "employee_db")

	fmt.Println("\n--- API settings ---")
	serverPort := ask(reader, "Server port", "8080")
	adminToken := askRequired(reader, "Choose an ADMIN token (can read and change data)")
	userToken := askRequired(reader, "Choose a USER token (read only)")
	if adminToken == userToken {
		fmt.Println("The two tokens have to be different.")
		return
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not set up the database connection:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Could not reach PostgreSQL:", err)
		return
	}

	// dependency injection: build from the bottom up and pass each piece upwards
	repo := repository.NewPostgresRepository(db)
	if err := repo.Init(); err != nil {
		fmt.Println("Could not create the table:", err)
		return
	}
	svc := service.NewEmployeeService(repo)
	employeeHandler := handler.NewEmployeeHandler(svc)
	healthHandler := handler.NewHealthHandler(svc.Ready)
	auth := middleware.NewAuth(adminToken, userToken)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler.Health)
	mux.HandleFunc("/ready", healthHandler.Ready)
	mux.Handle("/employee", auth.Authenticate(auth.Authorize(http.HandlerFunc(employeeHandler.Employees))))
	mux.Handle("/employee/", auth.Authenticate(auth.Authorize(http.HandlerFunc(employeeHandler.EmployeeByID))))

	// request ID goes outermost so the logger can print it
	app := middleware.RequestID(middleware.Logging(mux))

	log.Println("Employee API listening on port", serverPort)
	if err := http.ListenAndServe(":"+serverPort, app); err != nil {
		log.Fatal(err)
	}
}
