package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type apiError struct {
	Status  int
	Message string
	Fields  map[string]string
}

func (e *apiError) Error() string { return e.Message }

type errorBody struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

var (
	errNotFound  = &apiError{Status: http.StatusNotFound, Message: "employee not found"}
	errBadID     = &apiError{Status: http.StatusBadRequest, Message: "invalid employee id"}
	errBadBody   = &apiError{Status: http.StatusBadRequest, Message: "invalid JSON body"}
	errMethod    = &apiError{Status: http.StatusMethodNotAllowed, Message: "method not allowed"}
	errDuplicate = &apiError{Status: http.StatusConflict, Message: "email already exists"}
)

type Store struct {
	mu        sync.RWMutex
	employees []Employee
	nextID    int
}

func NewStore() *Store {
	return &Store{
		nextID: 4,
		employees: []Employee{
			{ID: 1, Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000},
			{ID: 2, Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000},
			{ID: 3, Name: "Sara Khan", Email: "sara@example.com", Department: "Engineering", Salary: 91000},
		},
	}
}

func (s *Store) List() []Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Employee, len(s.employees))
	copy(out, s.employees)
	return out
}

func (s *Store) Get(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, e := range s.employees {
		if e.ID == id {
			return e, nil
		}
	}
	return Employee{}, errNotFound
}

func (s *Store) emailTaken(email string, exceptID int) bool {
	for _, e := range s.employees {
		if e.ID != exceptID && strings.EqualFold(e.Email, email) {
			return true
		}
	}
	return false
}

func (s *Store) Create(e Employee) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.emailTaken(e.Email, 0) {
		return Employee{}, errDuplicate
	}
	e.ID = s.nextID
	s.nextID++
	s.employees = append(s.employees, e)
	return e, nil
}

func (s *Store) Update(id int, e Employee) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.employees {
		if s.employees[i].ID == id {
			if s.emailTaken(e.Email, id) {
				return Employee{}, errDuplicate
			}
			e.ID = id
			s.employees[i] = e
			return e, nil
		}
	}
	return Employee{}, errNotFound
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.employees {
		if s.employees[i].ID == id {
			s.employees = append(s.employees[:i], s.employees[i+1:]...)
			return nil
		}
	}
	return errNotFound
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func handle(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err == nil {
			return
		}
		var apiErr *apiError
		if errors.As(err, &apiErr) {
			writeJSON(w, apiErr.Status, errorBody{Error: apiErr.Message, Fields: apiErr.Fields})
			return
		}
		log.Println("internal error:", err)
		writeJSON(w, http.StatusInternalServerError, errorBody{Error: "internal server error"})
	}
}

func decodeBody(w http.ResponseWriter, r *http.Request, dst any) error {
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		return &apiError{Status: http.StatusUnsupportedMediaType, Message: "Content-Type must be application/json"}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		var syntaxErr *json.SyntaxError
		var typeErr *json.UnmarshalTypeError
		var maxErr *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxErr):
			return &apiError{Status: http.StatusBadRequest, Message: fmt.Sprintf("malformed JSON at position %d", syntaxErr.Offset)}
		case errors.Is(err, io.ErrUnexpectedEOF):
			return &apiError{Status: http.StatusBadRequest, Message: "malformed JSON"}
		case errors.As(err, &typeErr):
			return &apiError{Status: http.StatusBadRequest, Message: fmt.Sprintf("field %q must be of type %s", typeErr.Field, typeErr.Type)}
		case errors.Is(err, io.EOF):
			return &apiError{Status: http.StatusBadRequest, Message: "request body must not be empty"}
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return &apiError{Status: http.StatusBadRequest, Message: "unknown field " + field}
		case errors.As(err, &maxErr):
			return &apiError{Status: http.StatusRequestEntityTooLarge, Message: "request body too large"}
		default:
			return errBadBody
		}
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return &apiError{Status: http.StatusBadRequest, Message: "body must contain a single JSON object"}
	}
	return nil
}

var allowedDepartments = map[string]bool{
	"Engineering": true,
	"Finance":     true,
	"HR":          true,
	"Marketing":   true,
	"Sales":       true,
}

func validateEmployee(e *Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.ToLower(strings.TrimSpace(e.Email))
	e.Department = strings.TrimSpace(e.Department)

	fields := map[string]string{}

	if n := utf8.RuneCountInString(e.Name); n < 2 || n > 100 {
		fields["name"] = "must be between 2 and 100 characters"
	}

	addr, err := mail.ParseAddress(e.Email)
	if err != nil || addr.Address != e.Email {
		fields["email"] = "must be a valid email address"
	}

	if !allowedDepartments[e.Department] {
		fields["department"] = "must be one of: Engineering, Finance, HR, Marketing, Sales"
	}

	if e.Salary <= 0 || e.Salary > 10000000 {
		fields["salary"] = "must be greater than 0 and at most 10000000"
	}

	if len(fields) > 0 {
		return &apiError{Status: http.StatusUnprocessableEntity, Message: "validation failed", Fields: fields}
	}
	return nil
}

type API struct {
	store *Store
}

func parseID(path, prefix string) (int, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(path, prefix))
	if err != nil || id < 1 {
		return 0, errBadID
	}
	return id, nil
}

func (a *API) collection(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, a.store.List())
		return nil
	case http.MethodPost:
		var e Employee
		if err := decodeBody(w, r, &e); err != nil {
			return err
		}
		if err := validateEmployee(&e); err != nil {
			return err
		}
		created, err := a.store.Create(e)
		if err != nil {
			return err
		}
		w.Header().Set("Location", "/employee/"+strconv.Itoa(created.ID))
		writeJSON(w, http.StatusCreated, created)
		return nil
	default:
		w.Header().Set("Allow", "GET, POST")
		return errMethod
	}
}

func (a *API) item(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r.URL.Path, "/employee/")
	if err != nil {
		return err
	}

	switch r.Method {
	case http.MethodGet:
		e, err := a.store.Get(id)
		if err != nil {
			return err
		}
		writeJSON(w, http.StatusOK, e)
		return nil
	case http.MethodPut:
		var e Employee
		if err := decodeBody(w, r, &e); err != nil {
			return err
		}
		if err := validateEmployee(&e); err != nil {
			return err
		}
		updated, err := a.store.Update(id, e)
		if err != nil {
			return err
		}
		writeJSON(w, http.StatusOK, updated)
		return nil
	case http.MethodDelete:
		if err := a.store.Delete(id); err != nil {
			return err
		}
		w.WriteHeader(http.StatusNoContent)
		return nil
	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		return errMethod
	}
}

func main() {
	api := &API{store: NewStore()}

	mux := http.NewServeMux()
	mux.HandleFunc("/employee", handle(api.collection))
	mux.HandleFunc("/employee/", handle(api.item))

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
