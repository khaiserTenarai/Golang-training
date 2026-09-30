package employee

import (
	"errors"
	"math"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewManager()
	e, err := m.Add("  jane   doe ", "engineering", 50000, 2)
	if err != nil {
		t.Fatal(err)
	}
	if e.ID != 101 || e.Name != "jane doe" || e.Department != "Engineering" {
		t.Fatalf("unexpected: %+v", e)
	}
	got, err := m.Get(101)
	if err != nil || got != e {
		t.Fatalf("Get: %+v %v", got, err)
	}
}

func TestValidation(t *testing.T) {
	m := NewManager()
	bad := []struct {
		name, dept string
		sal        float64
		exp        int
	}{
		{"", "HR", 1, 1}, {"A", "Sales", 1, 1}, {"A", "HR", 0, 1}, {"A", "HR", 1, -1},
	}
	for _, b := range bad {
		if _, err := m.Add(b.name, b.dept, b.sal, b.exp); !errors.Is(err, ErrInvalid) {
			t.Errorf("%+v: want ErrInvalid, got %v", b, err)
		}
	}
}

func TestDeleteRebuildsIndex(t *testing.T) {
	m := NewSeeded() // IDs 101..105
	if err := m.Delete(102); err != nil {
		t.Fatal(err)
	}
	// Every later record moved left; lookups must still be correct.
	for _, id := range []int{101, 103, 104, 105} {
		e, err := m.Get(id)
		if err != nil || e.ID != id {
			t.Fatalf("after delete, Get(%d) = %+v, %v", id, e, err)
		}
	}
	if _, err := m.Get(102); !errors.Is(err, ErrNotFound) {
		t.Fatal("102 should be gone")
	}
	if m.Count() != 4 {
		t.Fatalf("count = %d", m.Count())
	}
}

func TestUpdate(t *testing.T) {
	m := NewSeeded()
	e, err := m.Update(102, "operations", 52000)
	if err != nil || e.Department != "Operations" || e.Salary != 52000 {
		t.Fatalf("%+v %v", e, err)
	}
	e, _ = m.Update(102, "", 0) // no change
	if e.Department != "Operations" || e.Salary != 52000 {
		t.Fatal("empty update changed data")
	}
	if _, err := m.Update(999, "", 1); !errors.Is(err, ErrNotFound) {
		t.Fatal("want not found")
	}
}

func TestSearchAndDept(t *testing.T) {
	m := NewSeeded()
	if r := m.SearchByName("RAO"); len(r) != 1 || r[0].Name != "Anita Rao" {
		t.Fatalf("search: %+v", r)
	}
	if r := m.ByDepartment("engineering"); len(r) != 2 {
		t.Fatalf("dept: %+v", r)
	}
}

func TestRaise(t *testing.T) {
	m := NewSeeded()
	n, err := m.GiveRaise("Engineering", 10)
	if err != nil || n != 2 {
		t.Fatalf("n=%d err=%v", n, err)
	}
	e, _ := m.Get(101)
	if math.Abs(e.Salary-99000) > 0.01 {
		t.Fatalf("salary %v", e.Salary)
	}
	if n, _ := m.GiveRaise("all", 5); n != 5 {
		t.Fatalf("all: %d", n)
	}
}

func TestListSortAndStats(t *testing.T) {
	m := NewSeeded()
	l := m.List(BySalaryDesc)
	if l[0].Name != "Anita Rao" || l[len(l)-1].Name != "Rahul Mehta" {
		t.Fatalf("sort: %+v", l)
	}
	if m.List(ByID)[0].ID != 101 {
		t.Fatal("original order changed")
	}
	s := m.Summary()
	if s.Count != 5 || s.Total != 320000 || s.Avg != 64000 || s.Highest.ID != 101 || s.Lowest.ID != 102 {
		t.Fatalf("stats %+v", s)
	}
	r := m.DepartmentReport()
	if len(r) != 4 || r[0].Department != "Engineering" || r[0].Count != 2 {
		t.Fatalf("report %+v", r)
	}
}
