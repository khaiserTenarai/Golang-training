package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// API holds everything the handlers need - the config, the secret, the
// logger and the (in-memory) data.
type API struct {
	cfg       Config
	secrets   Secrets
	log       *Logger
	mu        sync.Mutex
	employees []Employee
	nextID    int
}

func NewAPI(cfg Config, secrets Secrets, logger *Logger) *API {
	return &API{cfg: cfg, secrets: secrets, log: logger, employees: []Employee{}, nextID: 1}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Routes builds the router. The bonus endpoint and search are only
// there if their feature flags are on.
func (a *API) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", a.health)
	mux.HandleFunc("/employee", a.employeesHandler)
	mux.HandleFunc("/employee/", a.employeeByID)
	return a.logRequests(mux)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func (a *API) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)
		a.log.Infof("%s %s -> %d (%v)", r.Method, r.URL.Path, recorder.status, time.Since(start))
	})
}

// requireKey protects the endpoints that change data.
func (a *API) requireKey(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("X-API-Key") != a.secrets.APIKey.Reveal() {
		a.log.Warnf("rejected %s %s: missing or wrong API key", r.Method, r.URL.Path)
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing or wrong X-API-Key header"})
		return false
	}
	return true
}

func (a *API) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status": "ok",
		"env":    a.cfg.Env,
		"features": map[string]bool{
			"bonus":  a.cfg.FeatureBonus,
			"search": a.cfg.FeatureSearch,
		},
	})
}

// employeesHandler handles /employee (list and create)
func (a *API) employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.mu.Lock()
		defer a.mu.Unlock()

		results := a.employees
		if search := strings.ToLower(r.URL.Query().Get("search")); search != "" {
			if !a.cfg.FeatureSearch {
				writeJSON(w, http.StatusNotImplemented, map[string]string{"error": "search is switched off (FEATURE_SEARCH)"})
				return
			}
			results = []Employee{}
			for _, e := range a.employees {
				if strings.Contains(strings.ToLower(e.Name), search) {
					results = append(results, e)
				}
			}
		}
		writeJSON(w, http.StatusOK, results)

	case http.MethodPost:
		if !a.requireKey(w, r) {
			return
		}
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}
		if strings.TrimSpace(e.Name) == "" || e.Salary <= 0 {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name and a salary above 0 are required"})
			return
		}

		a.mu.Lock()
		e.ID = a.nextID
		a.nextID++
		a.employees = append(a.employees, e)
		a.mu.Unlock()

		a.log.Debugf("created employee %d (%s)", e.ID, e.Name)
		writeJSON(w, http.StatusCreated, e)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

// employeeByID handles /employee/{id}, /employee/{id}/bonus
func (a *API) employeeByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(strings.TrimPrefix(r.URL.Path, "/employee/"), "/"), "/")
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a number"})
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	idx := -1
	for i, e := range a.employees {
		if e.ID == id {
			idx = i
		}
	}
	if idx == -1 {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "employee not found"})
		return
	}

	// /employee/{id}/bonus - behind a feature flag
	if len(parts) == 2 && parts[1] == "bonus" {
		if !a.cfg.FeatureBonus {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, http.StatusOK, map[string]float64{"bonus": a.employees[idx].Salary * 0.10})
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, a.employees[idx])

	case http.MethodDelete:
		if !a.requireKey(w, r) {
			return
		}
		a.employees = append(a.employees[:idx], a.employees[idx+1:]...)
		a.log.Infof("deleted employee %d", id)
		w.WriteHeader(http.StatusNoContent)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}
