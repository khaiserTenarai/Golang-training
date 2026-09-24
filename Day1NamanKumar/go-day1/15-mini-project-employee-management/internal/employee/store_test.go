package employee

import (
	"errors"
	"testing"
)

func TestStoreLifecycle(t *testing.T) {
	s := NewStore()

	a, err := s.Add(Employee{Name: " Asha Rao ", Department: "Engineering", Salary: 100})
	if err != nil || a.ID != 1 || a.Name != "Asha Rao" {
		t.Fatalf("Add: got %+v, %v", a, err)
	}
	b, _ := s.Add(Employee{Name: "Vikram", Department: "Finance", Salary: 200})
	if b.ID != 2 {
		t.Fatalf("expected ID 2, got %d", b.ID)
	}

	if got, err := s.GetByID(1); err != nil || got.Name != "Asha Rao" {
		t.Fatalf("GetByID: %+v, %v", got, err)
	}
	if r := s.SearchByName("ASHA"); len(r) != 1 {
		t.Fatalf("SearchByName: want 1, got %d", len(r))
	}
	if l := s.List(); len(l) != 2 || l[0].ID != 1 {
		t.Fatalf("List: %+v", l)
	}
	if err := s.Delete(1); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := s.GetByID(1); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
	if err := s.Delete(99); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestAddValidation(t *testing.T) {
	s := NewStore()
	cases := map[string]struct {
		e    Employee
		want error
	}{
		"empty name":  {Employee{Department: "X", Salary: 1}, ErrInvalidName},
		"empty dept":  {Employee{Name: "A", Salary: 1}, ErrInvalidDept},
		"zero salary": {Employee{Name: "A", Department: "X"}, ErrInvalidSalry},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := s.Add(c.e); !errors.Is(err, c.want) {
				t.Errorf("got %v, want %v", err, c.want)
			}
		})
	}
}
