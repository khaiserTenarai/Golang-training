// Day 12 - Task 14: Create /ready endpoint
//
// /health says "the process is alive". /ready says "I'm actually able
// to serve traffic right now" - which usually means the things it
// depends on (a database, for example) are reachable. Here you give
// the address of one dependency, like localhost:5432 for PostgreSQL.
// /ready tries to connect to it on every call:
//   200 - reachable, ready for traffic
//   503 - not reachable, don't send traffic yet
// Start it with PostgreSQL running, call /ready, stop PostgreSQL, call
// it again, and watch it flip.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Address of the dependency to check (press Enter for localhost:5432): ")
	addrText, _ := reader.ReadString('\n')
	dependency := strings.TrimSpace(addrText)
	if dependency == "" {
		dependency = "localhost:5432"
	}

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		conn, err := net.DialTimeout("tcp", dependency, 2*time.Second)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "not ready",
				"reason": "can't reach " + dependency,
			})
			return
		}
		conn.Close()

		json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
	})

	fmt.Println("Server running on port", port, "- checking", dependency)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
