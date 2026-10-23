package saga

import (
    "context"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "employee-microservices/department"
    "employee-microservices/employee"
)

// Controller demonstrates an in-process orchestration Saga. The reserve
// functions below are intentionally simple placeholders for real operations.
type Controller struct {
    employees *employee.Service
    departments *department.Service
}
func NewController(e *employee.Service, d *department.Service) *Controller {
    return &Controller{employees:e, departments:d}
}

type SagaInput struct {
    EmployeeName string `json:"employee_name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    DepartmentID int64 `json:"department_id" binding:"required,gt=0"`
    FailDepartment bool `json:"fail_department"`
}

func (s *Controller) Run(c *gin.Context) {
    var input SagaInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
        return
    }
    requestCtx, cancel := contextWithTimeout(c, 5*time.Second)
    defer cancel()

    // Reserve employee resource: create employee record.
    created, err := s.employees.Create(requestCtx, employee.CreateInput{
        Name:input.EmployeeName, Email:input.Email, DepartmentID:input.DepartmentID,
    })
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"saga_status":"FAILED", "step":"reserve_employee", "error":err.Error()})
        return
    }

    // Reserve department resource: verify it again to demonstrate a failure
    // branch. In a real system, use an actual reservation operation.
    if input.FailDepartment {
        c.JSON(http.StatusConflict, gin.H{
            "saga_status":"COMPENSATION_REQUIRED",
            "message":"department reservation failed; employee release must be implemented as a real compensating operation",
            "employee_id":created.ID,
            "department_id":strconv.FormatInt(input.DepartmentID, 10),
        })
        return
    }
    if _, err := s.departments.Get(requestCtx, input.DepartmentID); err != nil {
        c.JSON(http.StatusConflict, gin.H{
            "saga_status":"COMPENSATION_REQUIRED",
            "message":"department step failed; implement a real employee compensation operation",
            "employee_id":created.ID,
        })
        return
    }
    c.JSON(http.StatusCreated, gin.H{"saga_status":"SUCCESS", "employee":created})
}

func contextWithTimeout(c *gin.Context, duration time.Duration) (context.Context, context.CancelFunc) {
    return context.WithTimeout(c.Request.Context(), duration)
}
