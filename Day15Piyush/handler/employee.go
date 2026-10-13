// Tasks 11, 12, 13: Employee CRUD + Validation + Pagination + File Upload

package handler

import (
	"fmt"
	"math"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"advanced_gin/models"

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
		{Name: "Diana", Email: "diana@test.com", Age: 27, Department: "HR", Salary: 50000},
	}
	for _, e := range seed {
		e.ID = h.next
		e.CreatedAt = time.Now()
		h.data[e.ID] = e
		h.next++
	}
	return h
}

// Task 12: Pagination + filtering + sorting
func (h *EmployeeHandler) GetAll(c *gin.Context) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	name := c.Query("name")
	dept := c.Query("department")
	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}

	var list []models.Employee
	for _, e := range h.data {
		if name != "" && !strings.Contains(strings.ToLower(e.Name), strings.ToLower(name)) {
			continue
		}
		if dept != "" && !strings.EqualFold(e.Department, dept) {
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
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, models.PaginatedResponse{
		Success: true, Data: list[start:end], Page: page,
		PageSize: pageSize, Total: total, TotalPages: totalPages,
	})
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.mu.RLock()
	emp, ok := h.data[id]
	h.mu.RUnlock()
	if !ok {
		c.JSON(http.StatusNotFound, models.APIResponse{Success: false, Message: "Not found"})
		return
	}
	c.JSON(http.StatusOK, models.APIResponse{Success: true, Data: emp})
}

// Task 11: Request validation via Gin binding tags
func (h *EmployeeHandler) Create(c *gin.Context) {
	var emp models.Employee
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

// Task 13: File Upload
func (h *EmployeeHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{Success: false, Message: "File required: " + err.Error()})
		return
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := map[string]bool{".csv": true, ".txt": true, ".json": true, ".pdf": true, ".jpg": true, ".png": true}
	if !allowed[ext] {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false, Message: "File type not allowed. Allowed: csv, txt, json, pdf, jpg, png",
		})
		return
	}

	// Validate size (max 5MB)
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, models.APIResponse{Success: false, Message: "File too large (max 5MB)"})
		return
	}

	savePath := fmt.Sprintf("./uploads/%d_%s", time.Now().UnixNano(), file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{Success: false, Message: "Failed to save: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true, Message: "File uploaded",
		Data: gin.H{"filename": file.Filename, "size": file.Size, "path": savePath},
	})
}
