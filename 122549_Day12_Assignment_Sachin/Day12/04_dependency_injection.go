package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

// Store is what the service depends on - any type with these two methods will do.
type Store interface {
	Add(e Employee) Employee
	All() []Employee
}

// ---------- storage option 1: a slice ----------

type sliceStore struct {
	mu    sync.Mutex
	items []Employee
}

func (s *sliceStore) Add(e Employee) Employee {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.ID = len(s.items) + 1
	s.items = append(s.items, e)
	return e
}

func (s *sliceStore) All() []Employee {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]Employee, len(s.items))
	copy(result, s.items)
	return result
}

// ---------- storage option 2: a map ----------

type mapStore struct {
	mu    sync.Mutex
	items map[int]Employee
}

func (m *mapStore) Add(e Employee) Employee {
	m.mu.Lock()
	defer m.mu.Unlock()
	e.ID = len(m.items) + 1
	m.items[e.ID] = e
	return e
}

func (m *mapStore) All() []Employee {
	m.mu.Lock()
	defer m.mu.Unlock()
	result := make([]Employee, 0, len(m.items))
	for _, e := range m.items {
		result = append(result, e)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

// ---------- service ----------

type EmployeeService struct {
	store Store
}

func NewEmployeeService(store Store) *EmployeeService {
	return &EmployeeService{store: store}
}

func (s *EmployeeService) Create(e Employee) Employee {
	return s.store.Add(e)
}

func (s *EmployeeService) List() []Employee {
	return s.store.All()
}

// ---------- handler ----------

type EmployeeHandler struct {
	service *EmployeeService
}

func NewEmployeeHandler(service *EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.service.List())

	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(h.service.Create(e))

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Which storage should be used? 1 = slice, 2 = map: ")
	choice, _ := reader.ReadString('\n')

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	// this is the "injection" - main picks the dependency and passes it in
	var store Store
	if strings.TrimSpace(choice) == "2" {
		store = &mapStore{items: map[int]Employee{}}
		fmt.Println("Using the map store.")
	} else {
		store = &sliceStore{}
		fmt.Println("Using the slice store.")
	}

	service := NewEmployeeService(store)
	handler := NewEmployeeHandler(service)

	http.HandleFunc("/employee", handler.Handle)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
