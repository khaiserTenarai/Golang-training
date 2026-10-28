package employee

import (
    "context"
    "encoding/json"
    "errors"
    "strings"

    "employee-microservices/department"
    "employee-microservices/eventbus"
)

type Service struct {
    repo *Repository
    departments *department.Service
    bus *eventbus.Bus
}
func NewService(repo *Repository, departments *department.Service, bus *eventbus.Bus) *Service {
    return &Service{repo: repo, departments: departments, bus: bus}
}
func (s *Service) Create(ctx context.Context, input CreateInput) (Employee, error) {
    input.Name = strings.TrimSpace(input.Name)
    input.Email = strings.TrimSpace(input.Email)
    if input.Name == "" || input.Email == "" || input.DepartmentID <= 0 {
        return Employee{}, errors.New("name, valid email, and department_id are required")
    }
    if _, err := s.departments.Get(ctx, input.DepartmentID); err != nil {
        return Employee{}, errors.New("department not found or department service unavailable")
    }
    e, err := s.repo.Create(ctx, Employee{
        Name: input.Name, Email: input.Email,
        DepartmentID: input.DepartmentID, Salary: input.Salary,
    })
    if err != nil { return Employee{}, err }

    payload, _ := json.Marshal(e)
    s.bus.Publish(eventbus.Event{Name: "EmployeeCreated", Data: string(payload)})
    return e, nil
}
func (s *Service) List(ctx context.Context) ([]Employee, error) {
    return s.repo.List(ctx)
}
