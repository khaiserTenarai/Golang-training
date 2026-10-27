package discovery

import (
	"encoding/json"
	"net/http"
)

// HealthHandler returns a handler for Consul health checks.
func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
		})
	}
}
