// Tasks 2–12: Employee Handlers using Gin
// Task 2  — GET routes
// Task 3  — POST routes
// Task 4  — PUT routes
// Task 5  — PATCH routes
// Task 6  — DELETE routes
// Task 8  — Path parameters (c.Param)
// Task 9  — Query parameters (c.Query, c.DefaultQuery)
// Task 10 — Bind JSON using Gin (c.ShouldBindJSON)
// Task 11 — Return JSON responses (c.JSON)
// Task 12 — Handle form parameters (c.ShouldBind with form tags)

package handler

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"gin_employee_api/models"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	mu   sync.RWMutex
	data map[int]models.Employee
	next int
}

func NewEmployeeHandler() *EmployeeHandler {
	h := &EmployeeHandler{data: make(map[int]models.Employee), next: 1}
	seed := []models.Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, Department: "Engineering", Salary: 60000},
		{Name: "Alice", Email: "alice@test.com", Age: 30, Department: "Design", Salary: 55000},
		{Name: "Bob", Email: "bob@test.com", Age: 28, Department: "Marketing", Salary: 45000},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, Department: "Engineering", Salary: 80000},
	}
	for _, e := range seed {
		e.ID = h.next
		e.CreatedAt = time.Now()
		h.data[e.ID] = e
		h.next++
	}
	return h
}

// Task 2: GET /api/employees — list all
// Task 9: Query parameters for filtering, sorting, pagination
// Task 11: Return JSON responses
func (h *EmployeeHandler) GetAll(c *gin.Context) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// Task 9: Query parameters
	name := c.Query("name")
	dept := c.Query("department")
	minSal, _ := strconv.ParseFloat(c.Query("min_salary"), 64)
	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var list []models.Employee
	for _, e := range h.data {
		if name != "" && !strings.Contains(strings.ToLower(e.Name), strings.ToLower(name)) {
			continue
		}
		if dept != "" && !strings.EqualFold(e.Department, dept) {
			continue
		}
		if minSal > 0 && e.Salary < minSal {
			continue
		}
		list = append(list, e)
	}

	desc := strings.ToLower(sortOrder) == "desc"
	sort.Slice(list, func(i, j int) bool {
		var less bool
		switch sortBy {
		case "name":
			less = list[i].Name < list[j].Name
		case "salary":
			less = list[i].Salary < list[j].Salary
		default:
			less = list[i].ID < list[j].ID
		}
		if desc {
			return !less
		}
		return less
	})

	total := len(list)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	// Task 11: JSON response
	c.JSON(http.StatusOK, models.PaginatedResponse{
		Success: true, Data: list[start:end], Total: total, Page: page, PageSize: pageSize,
	})
}

// Task 2: GET /api/employees/:id — detail
// Task 8: Path parameters
func (h *EmployeeHandler) GetByID(c *gin.Context) {
	// Task 8: Path parameter using c.Param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{Success: false, Message: "Invalid ID"})
		return
	}
	h.mu.RLock()
	emp, ok := h.data[id]
	h.mu.RUnlock()
	if !ok {
		c.JSON(http.StatusNotFound, models.APIResponse{Success: false, Message: "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, models.APIResponse{Success: true, Data: emp})
}

// Task 3: POST /api/employees — create
// Task 10: Bind JSON using Gin
func (h *EmployeeHandler) Create(c *gin.Context) {
	var emp models.Employee
	// Task 10: Gin JSON binding with validation
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false, Message: "Validation failed: " + err.Error(),
		})
		return
	}
	h.mu.Lock()
	emp.ID = h.next
	h.next++
	emp.CreatedAt = time.Now()
	h.data[emp.ID] = emp
	h.mu.Unlock()

	c.JSON(http.StatusCreated, models.APIResponse{Success: true, Message: "Created", Data: emp})
}

// Task 12: POST /api/employees/form — create from form data
func (h *EmployeeHandler) CreateFromForm(c *gin.Context) {
	var emp models.Employee
	// Task 12: Gin form binding
	if err := c.ShouldBind(&emp); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false, Message: "Form validation failed: " + err.Error(),
		})
		return
	}
	h.mu.Lock()
	emp.ID = h.next
	h.next++
	emp.CreatedAt = time.Now()
	h.data[emp.ID] = emp
	h.mu.Unlock()

	c.JSON(http.StatusCreated, models.APIResponse{Success: true, Message: "Created from form", Data: emp})
}

// Task 4: PUT /api/employees/:id — full update
func (h *EmployeeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	old, ok := h.data[id]
	if !ok {
		c.JSON(http.StatusNotFound, models.APIResponse{Success: false, Message: "Not found"})
		return
	}
	emp.ID = id
	emp.CreatedAt = old.CreatedAt
	h.data[id] = emp
	c.JSON(http.StatusOK, models.APIResponse{Success: true, Message: "Updated", Data: emp})
}

// Task 5: PATCH /api/employees/:id — partial update
func (h *EmployeeHandler) Patch(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var patch models.PatchEmployee
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	emp, ok := h.data[id]
	if !ok {
		c.JSON(http.StatusNotFound, models.APIResponse{Success: false, Message: "Not found"})
		return
	}
	// Apply only non-nil fields
	if patch.Name != nil {
		emp.Name = *patch.Name
	}
	if patch.Email != nil {
		emp.Email = *patch.Email
	}
	if patch.Age != nil {
		emp.Age = *patch.Age
	}
	if patch.Department != nil {
		emp.Department = *patch.Department
	}
	if patch.Salary != nil {
		emp.Salary = *patch.Salary
	}
	h.data[id] = emp
	c.JSON(http.StatusOK, models.APIResponse{Success: true, Message: "Patched", Data: emp})
}

// Task 6: DELETE /api/employees/:id
func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.data[id]; !ok {
		c.JSON(http.StatusNotFound, models.APIResponse{Success: false, Message: "Not found"})
		return
	}
	delete(h.data, id)
	c.JSON(http.StatusOK, models.APIResponse{Success: true, Message: "Deleted"})
}
