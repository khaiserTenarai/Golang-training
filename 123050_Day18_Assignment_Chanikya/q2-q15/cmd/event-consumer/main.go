package main

import (
	"log"
	"net/http"
	"time"

	"employee-management-microservices/events/consumer"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/events", consumer.Handler)

	server := &http.Server{
		Addr:              ":8084",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Event receiver running on port 8084")
	log.Fatal(server.ListenAndServe())
}
