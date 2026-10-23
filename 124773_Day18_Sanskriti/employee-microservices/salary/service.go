package salary

import (
    "context"
    "encoding/json"
    "log"
)

type Service struct { repo *Repository }
func NewService(repo *Repository) *Service { return &Service{repo:repo} }
func (s *Service) ListByEmployee(ctx context.Context, employeeID int64) ([]Salary, error) {
    return s.repo.ListByEmployee(ctx, employeeID)
}

// HandleEmployeeCreated is a demonstration event handler. Real salary rules
// should decide the amount and avoid duplicate processing using event IDs.
func (s *Service) HandleEmployeeCreated(payload string) {
    var employee struct {
        ID int64 `json:"id"`
        Salary float64 `json:"salary"`
    }
    if err := json.Unmarshal([]byte(payload), &employee); err != nil {
        log.Printf("salary consumer: invalid EmployeeCreated payload: %v", err)
        return
    }
    if employee.ID == 0 || employee.Salary <= 0 {
        log.Printf("salary consumer: received EmployeeCreated for employee %d; no salary row created (amount not provided)", employee.ID)
        return
    }
    if err := s.repo.Create(context.Background(), employee.ID, employee.Salary); err != nil {
        log.Printf("salary consumer: could not create salary: %v", err)
        return
    }
    log.Printf("salary consumer: initialized salary for employee %d", employee.ID)
}
