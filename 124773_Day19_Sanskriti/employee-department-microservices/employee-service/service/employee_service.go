package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"employee-department-microservices/employee-service/model"
	"employee-department-microservices/employee-service/repository"
	"employee-department-microservices/shared/dto"
	"employee-department-microservices/shared/utility"
)

type EmployeeService struct {
	repo          *repository.EmployeeRepository
	client        *http.Client
	departmentURL string
	retryCount    int
	breaker       *utility.CircuitBreaker
}

func NewEmployeeService(repo *repository.EmployeeRepository, client *http.Client, departmentURL string, retries int, breaker *utility.CircuitBreaker) *EmployeeService {
	return &EmployeeService{repo: repo, client: client, departmentURL: strings.TrimRight(departmentURL, "/"), retryCount: retries, breaker: breaker}
}
func (s *EmployeeService) List() []model.Employee { return s.repo.List() }
func (s *EmployeeService) Get(id int, correlationID string) (dto.EmployeeResponse, error) {
	e, err := s.repo.Get(id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	dept, err := s.getDepartment(e.DepartmentID, correlationID)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	return toResponse(e, &dept), nil
}
func (s *EmployeeService) Create(e model.Employee) model.Employee { return s.repo.Create(e) }
func (s *EmployeeService) Update(id int, e model.Employee) (model.Employee, error) {
	return s.repo.Update(id, e)
}
func (s *EmployeeService) Delete(id int) error { return s.repo.Delete(id) }
func (s *EmployeeService) CheckDepartment(correlationID string) error {
	if !s.breaker.Allow() {
		return utility.ErrCircuitOpen
	}
	_, err := s.requestDepartment("/health", correlationID)
	if err != nil {
		s.breaker.Failure()
		return err
	}
	s.breaker.Success()
	return nil
}
func (s *EmployeeService) GetDepartmentForValidation(id int, correlationID string) (dto.DepartmentResponse, error) {
	return s.getDepartment(id, correlationID)
}
func (s *EmployeeService) getDepartment(id int, correlationID string) (dto.DepartmentResponse, error) {
	if !s.breaker.Allow() {
		return dto.DepartmentResponse{}, utility.ErrCircuitOpen
	}
	var result dto.DepartmentResponse
	err := utility.Retry(s.retryCount, 100*time.Millisecond, func() error {
		data, err := s.requestDepartment(fmt.Sprintf("/departments/%d", id), correlationID)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, &result)
	})
	if err != nil {
		s.breaker.Failure()
		return result, err
	}
	s.breaker.Success()
	return result, nil
}
func (s *EmployeeService) requestDepartment(path, correlationID string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, s.departmentURL+path, nil)
	if err != nil {
		return nil, err
	}
	utility.ForwardCorrelation(req, correlationID)
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var payload json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("department service returned HTTP %d", resp.StatusCode)
	}
	return payload, nil
}
func toResponse(e model.Employee, d *dto.DepartmentResponse) dto.EmployeeResponse {
	return dto.EmployeeResponse{ID: e.ID, Name: e.Name, Email: e.Email, DepartmentID: e.DepartmentID, Department: d}
}
