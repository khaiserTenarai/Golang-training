package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"task15_employee_management_capstone/models"
	"task15_employee_management_capstone/service"

	"github.com/gorilla/mux"
)

type DepartmentController struct {
	svc service.DepartmentService
}

func NewDepartmentController(svc service.DepartmentService) *DepartmentController {
	return &DepartmentController{svc: svc}
}

func (c *DepartmentController) Create(w http.ResponseWriter, r *http.Request) {
	var dept models.Department
	if err := json.NewDecoder(r.Body).Decode(&dept); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	created, err := c.svc.CreateDepartment(dept)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "Department created", Data: created})
}

func (c *DepartmentController) GetAll(w http.ResponseWriter, r *http.Request) {
	depts, err := c.svc.GetAllDepartments()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error fetching departments"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Success", Data: depts})
}

func (c *DepartmentController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	dept, err := c.svc.GetDepartmentByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Department not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Success", Data: dept})
}

func (c *DepartmentController) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var dept models.Department
	if err := json.NewDecoder(r.Body).Decode(&dept); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	dept.ID = id
	updated, err := c.svc.UpdateDepartment(dept)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Department updated", Data: updated})
}

func (c *DepartmentController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.svc.DeleteDepartment(id); err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Department not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Department deleted"})
}
