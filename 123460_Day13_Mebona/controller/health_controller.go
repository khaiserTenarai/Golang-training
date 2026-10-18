package controller

import (
	"context"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

// Pinger is satisfied by service.EmployeeService.
type Pinger interface {
	Ping(ctx context.Context) error
}

type HealthController struct {
	ready   atomic.Bool
	pinger  Pinger
	summary func() map[string]any
}

func NewHealthController(p Pinger, summary func() map[string]any) *HealthController {
	return &HealthController{pinger: p, summary: summary}
}

// SetReady is flipped to true after startup and to false at the START of shutdown,
// so a load balancer stops sending new traffic while in-flight requests finish.
func (h *HealthController) SetReady(v bool) { h.ready.Store(v) }

// Live: "the process is running".
func (h *HealthController) Live(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Ready: "send me traffic".
func (h *HealthController) Ready(w http.ResponseWriter, r *http.Request) {
	if !h.ready.Load() {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "not ready"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := h.pinger.Ping(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "repository unavailable"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

// Config shows the active configuration with secrets redacted (not registered in production).
func (h *HealthController) Config(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.summary())
}

// Slow waits for ?seconds=N (max 30). Use it to watch graceful shutdown finish an in-flight request.
// Not registered in production.
func (h *HealthController) Slow(w http.ResponseWriter, r *http.Request) {
	secs, err := strconv.Atoi(r.URL.Query().Get("seconds"))
	if err != nil || secs < 1 || secs > 30 {
		secs = 5
	}
	select {
	case <-time.After(time.Duration(secs) * time.Second):
		writeJSON(w, http.StatusOK, map[string]any{"status": "finished", "waited": secs})
	case <-r.Context().Done(): // cancelled: client left, or the server was force-closed
		return
	}
}
