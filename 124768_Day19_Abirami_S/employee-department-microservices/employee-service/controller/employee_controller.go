package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-department-microservices/employee-service/model"
	"employee-department-microservices/employee-service/repository"
	"employee-department-microservices/employee-service/service"
	"employee-department-microservices/shared/dto"
	"employee-department-microservices/shared/utility"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct{ service *service.EmployeeService }

func NewEmployeeController(s *service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}
func (h *EmployeeController) Register(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"service": "employee", "status": "UP"}) })
	r.GET("/ready", func(c *gin.Context) {
		if err := h.service.CheckDepartment(utility.CorrelationID(c)); err != nil {
			utility.ErrorJSON(c, http.StatusServiceUnavailable, "department dependency unavailable")
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "READY"})
	})
	r.GET("/employees", h.list)
	r.GET("/employees/:id", h.get)
	r.POST("/employees", h.create)
	r.PUT("/employees/:id", h.update)
	r.DELETE("/employees/:id", h.delete)
}
func (h *EmployeeController) list(c *gin.Context) { c.JSON(http.StatusOK, h.service.List()) }
func (h *EmployeeController) get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	v, err := h.service.Get(id, utility.CorrelationID(c))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			utility.ErrorJSON(c, 404, err.Error())
		} else {
			utility.ErrorJSON(c, 503, err.Error())
		}
		return
	}
	c.JSON(200, v)
}
func (h *EmployeeController) create(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utility.ErrorJSON(c, 400, "invalid employee request: "+err.Error())
		return
	}
	// Confirm the referenced department exists before creating the employee.
	if _, err := h.service.GetDepartmentForValidation(req.DepartmentID, utility.CorrelationID(c)); err != nil {
		utility.ErrorJSON(c, 503, "department service unavailable or department not found")
		return
	}
	e := h.service.Create(model.Employee{Name: req.Name, Email: req.Email, DepartmentID: req.DepartmentID})
	c.JSON(http.StatusCreated, dto.EmployeeResponse{ID: e.ID, Name: e.Name, Email: e.Email, DepartmentID: e.DepartmentID})
}
func (h *EmployeeController) update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utility.ErrorJSON(c, 400, "invalid employee request: "+err.Error())
		return
	}
	if _, err := h.service.GetDepartmentForValidation(req.DepartmentID, utility.CorrelationID(c)); err != nil {
		utility.ErrorJSON(c, 503, "department service unavailable or department not found")
		return
	}
	e, err := h.service.Update(id, model.Employee{Name: req.Name, Email: req.Email, DepartmentID: req.DepartmentID})
	if errors.Is(err, repository.ErrNotFound) {
		utility.ErrorJSON(c, 404, err.Error())
		return
	}
	if err != nil {
		utility.ErrorJSON(c, 500, "could not update employee")
		return
	}
	c.JSON(200, dto.EmployeeResponse{ID: e.ID, Name: e.Name, Email: e.Email, DepartmentID: e.DepartmentID})
}
func (h *EmployeeController) delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	err := h.service.Delete(id)
	if errors.Is(err, repository.ErrNotFound) {
		utility.ErrorJSON(c, 404, err.Error())
		return
	}
	if err != nil {
		utility.ErrorJSON(c, 500, "could not delete employee")
		return
	}
	c.Status(204)
}
func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utility.ErrorJSON(c, 400, "id must be a positive integer")
		return 0, false
	}
	return id, true
}
