package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func flagEnabled(reader *bufio.Reader, key, description string) bool {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		fmt.Printf("Enable %s? (yes/no): ", description)
		text, _ := reader.ReadString('\n')
		value = strings.TrimSpace(text)
	}

	switch strings.ToLower(value) {
	case "1", "true", "yes", "y", "on":
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	reportEnabled := flagEnabled(reader, "FEATURE_REPORT", "the /report endpoint")
	newGreeting := flagEnabled(reader, "FEATURE_NEW_GREETING", "the new greeting message")

	fmt.Print("Port (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if newGreeting {
			fmt.Fprintln(w, "Welcome back! The new greeting is switched on.")
			return
		}
		fmt.Fprintln(w, "Hello from the employee API.")
	})

	// the route is only registered when the flag is on
	if reportEnabled {
		mux.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"employees": 42, "average_salary": 51250.5})
		})
	}

	fmt.Printf("\nFlags: report=%v new_greeting=%v\n", reportEnabled, newGreeting)
	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
