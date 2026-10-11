// Package controller contains Gin HTTP handlers.
package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
)

// EmployeeController handles employee requests.
type EmployeeController struct {
	service service.EmployeeService
}

// NewEmployeeController creates the controller.
func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}

// RegisterRoutes registers employee REST routes, route groups, static files, and form binding.
func (c *EmployeeController) RegisterRoutes(r *gin.Engine) {
	// Point 13: Serve static files
	r.Static("/static", "./static")

	// Point 12: Handle form parameters
	r.POST("/employees/form", c.CreateFromForm)

	// Point 7: Route groups (v1)
	v1 := r.Group("/api/v1")
	{
		// Points 2, 3, 4, 5, 6: REST verbs (GET, POST, PUT, PATCH, DELETE)
		v1.GET("/employees", c.GetAll)
		v1.POST("/employees", c.Create)
		v1.GET("/employees/:id", c.GetByID)
		v1.PUT("/employees/:id", c.Update)
		v1.PATCH("/employees/:id", c.PartialUpdate)
		v1.DELETE("/employees/:id", c.Delete)
	}
}

// Point 9: Query parameters handling
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	sort := ctx.Query("sort")
	e, err := c.service.GetAll(ctx.Request.Context(), sort)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, e)
}

// Point 10 & 11: Bind JSON & Return JSON responses
func (c *EmployeeController) Create(ctx *gin.Context) {
	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload or missing fields"})
		return
	}
	if err := c.service.Create(ctx.Request.Context(), &e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, e)
}

// Point 12: Handle form parameters (x-www-form-urlencoded or multipart/form-data)
func (c *EmployeeController) CreateFromForm(ctx *gin.Context) {
	salary, _ := strconv.ParseFloat(ctx.PostForm("salary"), 64)
	emp := model.Employee{
		Name:       ctx.PostForm("name"),
		Email:      ctx.PostForm("email"),
		Department: ctx.PostForm("department"),
		Salary:     salary,
	}

	if err := c.service.Create(ctx.Request.Context(), &emp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, emp)
}

// Point 8: Path parameters handling
func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "employee id must be numeric"})
		return
	}
	e, err := c.service.GetByID(ctx.Request.Context(), id)
	if repository.IsNotFound(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, e)
}

// Point 4: PUT Route
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "employee id must be numeric"})
		return
	}
	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
		return
	}
	u, err := c.service.Update(ctx.Request.Context(), id, &e)
	if repository.IsNotFound(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// Point 5: PATCH Route
func (c *EmployeeController) PartialUpdate(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "employee id must be numeric"})
		return
	}
	var patch model.PartialEmployee
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
		return
	}
	u, err := c.service.PartialUpdate(ctx.Request.Context(), id, &patch)
	if repository.IsNotFound(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// Point 6: DELETE Route
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "employee id must be numeric"})
		return
	}
	if err := c.service.Delete(ctx.Request.Context(), id); err != nil {
		if repository.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}