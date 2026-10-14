package controller

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"day-14-gin-employee-api/model"
	"day-14-gin-employee-api/repository"
	"day-14-gin-employee-api/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service *service.EmployeeService
}

func NewEmployeeController(s *service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

func (ctl *EmployeeController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	employees := api.Group("/employees")
	{
		//Task 2 to 8
		employees.POST("", ctl.Create)
		employees.GET("", ctl.GetAll)
		employees.GET("/:id", ctl.GetByID)
		employees.PUT("/:id", ctl.Update)
		employees.PATCH("/:id", ctl.Patch)
		employees.DELETE("/:id", ctl.Delete)
		employees.POST("/form", ctl.CreateFromForm)
	}
	r.GET("/hello/:name", ctl.Hello)
}

func (ctl *EmployeeController) Create(c *gin.Context) {
	var e model.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee JSON: " + err.Error()})
		return
	}
	if err := ctl.service.Create(c.Request.Context(), &e); err != nil {
		ctl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, e)
}

func (ctl *EmployeeController) GetAll(c *gin.Context) {
	employees, err := ctl.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch employees"})
		return
	}
	department := strings.TrimSpace(c.Query("department"))
	if department != "" {
		filtered := make([]model.Employee, 0)
		for _, e := range employees {
			if strings.EqualFold(e.Department, department) {
				filtered = append(filtered, e)
			}
		}
		employees = filtered
	}
	c.JSON(http.StatusOK, gin.H{"data": employees, "count": len(employees)})
}

func (ctl *EmployeeController) GetByID(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	e, err := ctl.service.GetByID(c.Request.Context(), id)
	if err != nil {
		ctl.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, e)
}

func (ctl *EmployeeController) Update(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	var e model.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee JSON: " + err.Error()})
		return
	}
	updated, err := ctl.service.Update(c.Request.Context(), id, &e)
	if err != nil {
		ctl.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (ctl *EmployeeController) Patch(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	var fields map[string]any
	if err := c.ShouldBindJSON(&fields); err != nil || len(fields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provide a non-empty JSON object"})
		return
	}
	updated, err := ctl.service.Patch(c.Request.Context(), id, fields)
	if err != nil {
		ctl.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (ctl *EmployeeController) Delete(c *gin.Context) {
	id, ok := getID(c)
	if !ok {
		return
	}
	if err := ctl.service.Delete(c.Request.Context(), id); err != nil {
		ctl.handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctl *EmployeeController) CreateFromForm(c *gin.Context) {
	e := model.Employee{
		Name:       c.PostForm("name"),
		Email:      c.PostForm("email"),
		Department: c.PostForm("department"),
	}
	if salary := c.PostForm("salary"); salary != "" {
		value, err := strconv.ParseFloat(salary, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "salary must be numeric"})
			return
		}
		e.Salary = value
	}
	if err := ctl.service.Create(c.Request.Context(), &e); err != nil {
		ctl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, e)
}

func (ctl *EmployeeController) Hello(c *gin.Context) {
	name := c.Param("name")
	language := c.DefaultQuery("lang", "en")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name, "language": language})
}

func getID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func (ctl *EmployeeController) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, repository.ErrEmailExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
