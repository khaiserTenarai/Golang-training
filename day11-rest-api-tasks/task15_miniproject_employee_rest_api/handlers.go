package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

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

type listResponse struct {
	Data []Employee `json:"data"`
	Meta listMeta   `json:"meta"`
}

type listMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

var (
	errBadID            = &apiError{Status: http.StatusBadRequest, Message: "invalid employee id"}
	errMethod           = &apiError{Status: http.StatusMethodNotAllowed, Message: "method not allowed"}
	errEmployeeNotFound = &apiError{Status: http.StatusNotFound, Message: "employee not found"}
	errDuplicate        = &apiError{Status: http.StatusConflict, Message: "email already exists"}
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func toAPIError(err error) *apiError {
	var apiErr *apiError
	switch {
	case errors.As(err, &apiErr):
		return apiErr
	case errors.Is(err, ErrNotFound):
		return errEmployeeNotFound
	case errors.Is(err, ErrDuplicateEmail):
		return errDuplicate
	}
	return nil
}

func handle(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err == nil {
			return
		}
		if apiErr := toAPIError(err); apiErr != nil {
			writeJSON(w, apiErr.Status, errorBody{Error: apiErr.Message, Fields: apiErr.Fields})
			return
		}
		log.Printf("internal error: %v", err)
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
			return &apiError{Status: http.StatusBadRequest, Message: "unknown field " + strings.TrimPrefix(err.Error(), "json: unknown field ")}
		case errors.As(err, &maxErr):
			return &apiError{Status: http.StatusRequestEntityTooLarge, Message: "request body too large"}
		default:
			return &apiError{Status: http.StatusBadRequest, Message: "invalid JSON body"}
		}
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return &apiError{Status: http.StatusBadRequest, Message: "body must contain a single JSON object"}
	}
	return nil
}

func decodeInput(w http.ResponseWriter, r *http.Request) (EmployeeInput, error) {
	var in EmployeeInput
	if err := decodeBody(w, r, &in); err != nil {
		return in, err
	}
	if fields := validateInput(&in); len(fields) > 0 {
		return in, &apiError{Status: http.StatusUnprocessableEntity, Message: "validation failed", Fields: fields}
	}
	return in, nil
}

func parseListParams(q url.Values) (ListParams, error) {
	badRequest := func(msg string) (ListParams, error) {
		return ListParams{}, &apiError{Status: http.StatusBadRequest, Message: msg}
	}

	p := ListParams{Page: 1, Limit: 10, Sort: "id", Order: "asc"}

	if v := q.Get("page"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			return badRequest("page must be a positive integer")
		}
		p.Page = n
	}

	if v := q.Get("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 || n > 100 {
			return badRequest("limit must be between 1 and 100")
		}
		p.Limit = n
	}

	if v := q.Get("sort"); v != "" {
		if _, ok := sortColumns[v]; !ok {
			return badRequest("sort must be one of: id, name, salary, department, created_at")
		}
		p.Sort = v
	}

	if v := strings.ToLower(q.Get("order")); v != "" {
		if v != "asc" && v != "desc" {
			return badRequest("order must be asc or desc")
		}
		p.Order = v
	}

	p.Department = strings.TrimSpace(q.Get("department"))
	p.Search = strings.TrimSpace(q.Get("search"))
	return p, nil
}

type API struct {
	repo Repository
}

func (a *API) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := a.repo.Ping(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "degraded", "database": "down"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "database": "up"})
}

func (a *API) collection(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		params, err := parseListParams(r.URL.Query())
		if err != nil {
			return err
		}
		employees, total, err := a.repo.List(r.Context(), params)
		if err != nil {
			return err
		}
		writeJSON(w, http.StatusOK, listResponse{
			Data: employees,
			Meta: listMeta{
				Page:       params.Page,
				Limit:      params.Limit,
				Total:      total,
				TotalPages: (total + params.Limit - 1) / params.Limit,
			},
		})
		return nil
	case http.MethodPost:
		in, err := decodeInput(w, r)
		if err != nil {
			return err
		}
		created, err := a.repo.Create(r.Context(), in)
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
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil || id < 1 {
		return errBadID
	}

	switch r.Method {
	case http.MethodGet:
		e, err := a.repo.GetByID(r.Context(), id)
		if err != nil {
			return err
		}
		writeJSON(w, http.StatusOK, e)
		return nil
	case http.MethodPut:
		in, err := decodeInput(w, r)
		if err != nil {
			return err
		}
		updated, err := a.repo.Update(r.Context(), id, in)
		if err != nil {
			return err
		}
		writeJSON(w, http.StatusOK, updated)
		return nil
	case http.MethodDelete:
		if err := a.repo.Delete(r.Context(), id); err != nil {
			return err
		}
		w.WriteHeader(http.StatusNoContent)
		return nil
	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		return errMethod
	}
}

type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

func (sr *statusRecorder) Write(b []byte) (int, error) {
	n, err := sr.ResponseWriter.Write(b)
	sr.bytes += n
	return n, err
}

var requestCounter uint64

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		id := atomic.AddUint64(&requestCounter, 1)

		w.Header().Set("X-Request-ID", fmt.Sprintf("req-%d", id))
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		log.Printf("[req-%d] %s %s %d %dB %s", id, r.Method, r.URL.RequestURI(), rec.status, rec.bytes, time.Since(start))
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic recovered: %v", rec)
				writeJSON(w, http.StatusInternalServerError, errorBody{Error: "internal server error"})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func NewRouter(repo Repository) http.Handler {
	api := &API{repo: repo}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", api.health)
	mux.HandleFunc("/employee", handle(api.collection))
	mux.HandleFunc("/employee/", handle(api.item))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusNotFound, errorBody{Error: "route not found"})
	})

	return loggingMiddleware(recoveryMiddleware(mux))
}
