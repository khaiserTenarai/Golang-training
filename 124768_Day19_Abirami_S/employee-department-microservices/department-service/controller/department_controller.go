package controller

import (
	"employee-department-microservices/department-service/model"
	"employee-department-microservices/department-service/repository"
	"employee-department-microservices/department-service/service"
	"employee-department-microservices/shared/dto"
	"employee-department-microservices/shared/utility"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DepartmentController struct{ s *service.DepartmentService }

func NewDepartmentController(s *service.DepartmentService) *DepartmentController {
	return &DepartmentController{s: s}
}
func (h *DepartmentController) Register(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "department", "status": "UP"}) })
	r.GET("/departments", h.list)
	r.GET("/departments/:id", h.get)
	r.POST("/departments", h.create)
	r.PUT("/departments/:id", h.update)
	r.DELETE("/departments/:id", h.delete)
}
func (h *DepartmentController) list(c *gin.Context) {
	items := h.s.List()
	out := make([]dto.DepartmentResponse, 0, len(items))
	for _, v := range items {
		out = append(out, respond(v))
	}
	c.JSON(200, out)
}
func (h *DepartmentController) get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	v, err := h.s.Get(id)
	if errors.Is(err, repository.ErrNotFound) {
		utility.ErrorJSON(c, 404, err.Error())
		return
	}
	if err != nil {
		utility.ErrorJSON(c, 500, "could not get department")
		return
	}
	c.JSON(200, respond(v))
}
func (h *DepartmentController) create(c *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utility.ErrorJSON(c, 400, "invalid department request: "+err.Error())
		return
	}
	v := h.s.Create(model.Department{Name: req.Name, Description: req.Description})
	c.JSON(http.StatusCreated, respond(v))
}
func (h *DepartmentController) update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utility.ErrorJSON(c, 400, "invalid department request: "+err.Error())
		return
	}
	v, err := h.s.Update(id, model.Department{Name: req.Name, Description: req.Description})
	if errors.Is(err, repository.ErrNotFound) {
		utility.ErrorJSON(c, 404, err.Error())
		return
	}
	if err != nil {
		utility.ErrorJSON(c, 500, "could not update department")
		return
	}
	c.JSON(200, respond(v))
}
func (h *DepartmentController) delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	err := h.s.Delete(id)
	if errors.Is(err, repository.ErrNotFound) {
		utility.ErrorJSON(c, 404, err.Error())
		return
	}
	if err != nil {
		utility.ErrorJSON(c, 500, "could not delete department")
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
func respond(v model.Department) dto.DepartmentResponse {
	return dto.DepartmentResponse{ID: v.ID, Name: v.Name, Description: v.Description}
}
