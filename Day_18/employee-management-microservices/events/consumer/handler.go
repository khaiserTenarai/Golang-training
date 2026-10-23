package consumer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Event is the event format received over HTTP.
type Event struct {
	EventID   string          `json:"eventId"`
	EventType string          `json:"eventType"`
	Payload   json.RawMessage `json:"payload"`
}

// Handler accepts events from the outbox worker.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var event Event
	if err := json.Unmarshal(body, &event); err != nil ||
		event.EventID == "" || event.EventType == "" {
		http.Error(w, "Invalid event", http.StatusBadRequest)
		return
	}

	// This example logs events; durable consumers should store event IDs
	// and ignore IDs that have already been processed.
	log.Printf("Received event %s: %s", event.EventID, event.EventType)

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte(`{"status":"accepted"}`))
}
