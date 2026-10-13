package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name of this service (shown in the health reply): ")
	nameText, _ := reader.ReadString('\n')
	serviceName := strings.TrimSpace(nameText)
	if serviceName == "" {
		serviceName = "employee-api"
	}

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	started := time.Now()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"service": serviceName,
			"uptime":  time.Since(started).Round(time.Second).String(),
		})
	})

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
