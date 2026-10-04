package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"task15_employee_management_capstone/models"
	"task15_employee_management_capstone/repository"
	"task15_employee_management_capstone/service"

	"github.com/gorilla/mux"
)

type EmployeeController struct {
	svc service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{svc: svc}
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	created, err := c.svc.CreateEmployee(emp)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "Employee created", Data: created})
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	search := q.Get("search")
	sortBy := q.Get("sort_by")
	sortOrder := q.Get("sort_order")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	employees, total, err := c.svc.GetAllEmployees(page, pageSize, search, sortBy, sortOrder)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error fetching employees"})
		return
	}

	sendJSON(w, http.StatusOK, models.PaginatedResponse{
		Message:    "Success",
		Data:       employees,
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		TotalPages: repository.TotalPages(total, pageSize),
	})
}

func (c *EmployeeController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	emp, err := c.svc.GetEmployeeByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Success", Data: emp})
}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	emp.ID = id
	updated, err := c.svc.UpdateEmployee(emp)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee updated", Data: updated})
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.svc.DeleteEmployee(id); err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee deleted"})
}

func (c *EmployeeController) UpdateSalary(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var req models.SalaryUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	if err := c.svc.UpdateSalary(id, req.NewSalary, req.Reason); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Salary updated via transaction"})
}

func (c *EmployeeController) SalaryHistory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	history, err := c.svc.GetSalaryHistory(id)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Salary history", Data: history})
}
