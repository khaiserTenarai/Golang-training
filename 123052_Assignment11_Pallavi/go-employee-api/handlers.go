package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type Server struct {
	store Store
}

func (s *Server) handleEmployees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getEmployees(w, r)
	case http.MethodPost:
		s.createEmployee(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.getEmployeeByID(w, r, id)
	case http.MethodPut:
		s.updateEmployee(w, r, id)
	case http.MethodDelete:
		s.deleteEmployee(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) getEmployees(w http.ResponseWriter, r *http.Request) {
	roleQuery := r.URL.Query().Get("role")

	employees, err := s.store.GetAllEmployees(roleQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (s *Server) getEmployeeByID(w http.ResponseWriter, r *http.Request, id int) {
	emp, err := s.store.GetEmployeeByID(id)
	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (s *Server) createEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if msg := e.Validate(); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if err := s.store.CreateEmployee(&e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (s *Server) updateEmployee(w http.ResponseWriter, r *http.Request, id int) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if msg := e.Validate(); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if err := s.store.UpdateEmployee(id, &e); err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func (s *Server) deleteEmployee(w http.ResponseWriter, r *http.Request, id int) {
	if err := s.store.DeleteEmployee(id); err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}