package employee

import (
	"errors"
	"sync"
	"testing"
)

func TestServiceAddGet(t *testing.T) {
	s := NewService()
	e := Employee{ID: 1, Name: "Asha", Email: "a@x.com", Age: 30, BaseSalary: 50000}

	if err := s.Add(e); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if err := s.Add(e); !errors.Is(err, ErrDuplicate) {
		t.Errorf("duplicate Add = %v; want ErrDuplicate", err)
	}
	got, err := s.Get(1)
	if err != nil || got.Name != "Asha" {
		t.Errorf("Get = %+v, %v", got, err)
	}
	if _, err := s.Get(99); !errors.Is(err, ErrNotFound) {
		t.Errorf("Get(99) = %v; want ErrNotFound", err)
	}
	if err := s.Add(Employee{}); err == nil {
		t.Error("expected validation error for empty employee")
	}
}

func TestServiceConcurrentAdd(t *testing.T) {
	s := NewService()
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = s.Add(Employee{ID: id, Name: "E", Email: "e@x.com", Age: 30, BaseSalary: 1000})
		}(i)
	}
	wg.Wait()
	if s.Count() != 100 {
		t.Errorf("Count = %d; want 100", s.Count())
	}
}