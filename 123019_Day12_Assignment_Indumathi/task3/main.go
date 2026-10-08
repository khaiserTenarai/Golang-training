package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

// Repository Interface abstracts storage layer
type Repository interface {
	FindByID(id int) (*Employee, error)
}

type memoryRepo struct{}

func (r *memoryRepo) FindByID(id int) (*Employee, error) {
	return &Employee{ID: id, Name: "Bob"}, nil
}

// Service contains domain logic, oblivious to SQL/NoSQL storage specifics
type Service struct {
	repo Repository
}

func (s *Service) GetFormattedName(id int) (string, error) {
	emp, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return "Employee: " + emp.Name, nil
}

func main() {
	repo := &memoryRepo{}
	svc := &Service{repo: repo}

	out, _ := svc.GetFormattedName(1)
	fmt.Println(out)
}