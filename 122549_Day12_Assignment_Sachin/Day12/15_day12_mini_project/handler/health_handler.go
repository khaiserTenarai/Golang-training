package handler

import (
	"net/http"
	"time"
)

type HealthHandler struct {
	readyCheck func() error
	started    time.Time
}

func NewHealthHandler(readyCheck func() error) *HealthHandler {
	return &HealthHandler{readyCheck: readyCheck, started: time.Now()}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"uptime": time.Since(h.started).Round(time.Second).String(),
	})
}

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
