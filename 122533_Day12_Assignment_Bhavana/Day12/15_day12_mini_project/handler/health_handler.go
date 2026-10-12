package handler

import (
	"net/http"
	"time"
)

type HealthHandler struct {
	readyCheck func() error
	started    time.Time
}

// readyCheck is whatever function tells us if our dependencies are OK.
func NewHealthHandler(readyCheck func() error) *HealthHandler {
	return &HealthHandler{readyCheck: readyCheck, started: time.Now()}
}

// Health: is the process alive?
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"uptime": time.Since(h.started).Round(time.Second).String(),
	})
}

// Ready: can we actually serve traffic right now (is the database reachable)?
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.readyCheck(); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{
			"status": "not ready",
			"reason": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}
