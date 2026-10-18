package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"employee-api/model"
	"employee-api/service"
)

type EmployeeController struct {
	svc          *service.EmployeeService
	logger       *slog.Logger
	exposeSalary bool // FEATURE_EXPOSE_SALARY
}

func NewEmployeeController(svc *service.EmployeeService, logger *slog.Logger, exposeSalary bool) *EmployeeController {
	return &EmployeeController{svc: svc, logger: logger, exposeSalary: exposeSalary}
}

// Register adds the routes. Reading is public; creating, changing and deleting need the API key.
func (c *EmployeeController) Register(mux *http.ServeMux, writeGuard func(http.Handler) http.Handler) {
	mux.HandleFunc("GET /employees", c.List)
	mux.HandleFunc("GET /employees/{id}", c.Get)
	mux.Handle("POST /employees", writeGuard(http.HandlerFunc(c.Create)))
	mux.Handle("PUT /employees/{id}", writeGuard(http.HandlerFunc(c.Update)))
	mux.Handle("DELETE /employees/{id}", writeGuard(http.HandlerFunc(c.Delete)))
}

// employeeResponse is what leaves the API. Salary is a pointer so it can disappear completely
// (omitempty) when the feature flag is off.
type employeeResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	Salary     *float64  `json:"salary,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (c *EmployeeController) toResponse(e model.Employee) employeeResponse {
	r := employeeResponse{ID: e.ID, Name: e.Name, Email: e.Email, Department: e.Department, CreatedAt: e.CreatedAt}
	if c.exposeSalary {
		salary := e.Salary
		r.Salary = &salary
	}
	return r
}

func (c *EmployeeController) List(w http.ResponseWriter, r *http.Request) {
	page, err := intQuery(r, "page")
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	size, err := intQuery(r, "pageSize")
	if err != nil {
		writeError(w, c.logger, err)
		return
	}

	// r.Context() flows into the service and then into the repository
	res, err := c.svc.List(r.Context(), service.ListQuery{Search: r.URL.Query().Get("search"), Page: page, PageSize: size})
	if err != nil {
		writeError(w, c.logger, err)
		return
	}

	items := make([]employeeResponse, 0, len(res.Items))
	for _, e := range res.Items {
		items = append(items, c.toResponse(e))
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"items": items, "total": res.Total, "page": res.Page, "pageSize": res.PageSize, "totalPages": res.TotalPages,
	})
}

func (c *EmployeeController) Get(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	e, err := c.svc.Get(r.Context(), id)
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, c.toResponse(e))
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var in model.EmployeeInput
	if err := decodeJSON(w, r, &in); err != nil {
		writeError(w, c.logger, err)
		return
	}
	e, err := c.svc.Create(r.Context(), in)
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/employees/%d", e.ID))
	writeJSON(w, http.StatusCreated, c.toResponse(e))
}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	var in model.EmployeeInput
	if err := decodeJSON(w, r, &in); err != nil {
		writeError(w, c.logger, err)
		return
	}
	e, err := c.svc.Update(r.Context(), id, in)
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	writeJSON(w, http.StatusOK, c.toResponse(e))
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, c.logger, err)
		return
	}
	if err := c.svc.Delete(r.Context(), id); err != nil {
		writeError(w, c.logger, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ---------- small helpers ----------

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return &model.ValidationError{Fields: map[string]string{"body": "invalid JSON: " + err.Error()}}
	}
	return nil
}

func pathID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		return 0, &model.ValidationError{Fields: map[string]string{"id": "must be a positive integer"}}
	}
	return id, nil
}

// intQuery returns 0 when the parameter is missing (the service then applies its default).
func intQuery(r *http.Request, name string) (int, error) {
	v := r.URL.Query().Get(name)
	if v == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, &model.ValidationError{Fields: map[string]string{name: "must be a whole number"}}
	}
	return n, nil
}
