package employee

type Employee struct {
    ID           int64   `json:"id"`
    Name         string  `json:"name"`
    Email        string  `json:"email"`
    DepartmentID int64   `json:"department_id"`
    Salary       float64 `json:"salary,omitempty"`
}

type CreateInput struct {
    Name         string  `json:"name" binding:"required"`
    Email        string  `json:"email" binding:"required,email"`
    DepartmentID int64   `json:"department_id" binding:"required,gt=0"`
    Salary       float64 `json:"salary"`
}
