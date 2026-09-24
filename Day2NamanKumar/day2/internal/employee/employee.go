// Package employee is an in-memory employee store for the Day 2 mini
// project. It combines two collections:
//
//   - a SLICE ([]Employee) that keeps records in insertion order, and
//   - a MAP (map[int]int) from employee ID to its position in the slice,
//     so lookups by ID are O(1) instead of scanning the whole slice.
package employee

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Errors returned by Manager methods.
var (
	ErrNotFound = errors.New("employee not found")
	ErrInvalid  = errors.New("invalid input")
)

// Departments lists the allowed department names.
var Departments = []string{"Engineering", "HR", "Finance", "Marketing", "Operations"}

// Employee is a single staff record.
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	Experience int // years
}

// Manager stores employees in a slice plus an ID index map.
type Manager struct {
	list   []Employee
	index  map[int]int // ID -> position in list
	nextID int
}

// NewManager returns an empty Manager whose first ID is 101.
func NewManager() *Manager {
	return &Manager{index: make(map[int]int), nextID: 101}
}

// NewSeeded returns a Manager with sample employees.
func NewSeeded() *Manager {
	m := NewManager()
	seed := []Employee{
		{Name: "Anita Rao", Department: "Engineering", Salary: 90000, Experience: 6},
		{Name: "Rahul Mehta", Department: "HR", Salary: 45000, Experience: 2},
		{Name: "Priya Nair", Department: "Finance", Salary: 60000, Experience: 4},
		{Name: "Arjun Singh", Department: "Engineering", Salary: 70000, Experience: 3},
		{Name: "Kavya Sharma", Department: "Marketing", Salary: 55000, Experience: 5},
	}
	for _, e := range seed {
		_, _ = m.Add(e.Name, e.Department, e.Salary, e.Experience)
	}
	return m
}

// Add validates the data, assigns the next ID and stores the employee.
func (m *Manager) Add(name, dept string, salary float64, exp int) (Employee, error) {
	name = strings.Join(strings.Fields(name), " ")
	d, ok := NormalizeDept(dept)
	switch {
	case name == "":
		return Employee{}, fmt.Errorf("%w: name is required", ErrInvalid)
	case !ok:
		return Employee{}, fmt.Errorf("%w: department must be one of %v", ErrInvalid, Departments)
	case salary <= 0:
		return Employee{}, fmt.Errorf("%w: salary must be positive", ErrInvalid)
	case exp < 0 || exp > 50:
		return Employee{}, fmt.Errorf("%w: experience must be 0-50 years", ErrInvalid)
	}
	e := Employee{ID: m.nextID, Name: name, Department: d, Salary: salary, Experience: exp}
	m.list = append(m.list, e)
	m.index[e.ID] = len(m.list) - 1
	m.nextID++
	return e, nil
}

// Get returns the employee with id using the map index (O(1)).
func (m *Manager) Get(id int) (Employee, error) {
	pos, ok := m.index[id]
	if !ok {
		return Employee{}, fmt.Errorf("%w: ID %d", ErrNotFound, id)
	}
	return m.list[pos], nil
}

// SearchByName returns employees whose name contains q (case-insensitive).
func (m *Manager) SearchByName(q string) []Employee {
	q = strings.ToLower(strings.TrimSpace(q))
	var out []Employee
	for _, e := range m.list {
		if q != "" && strings.Contains(strings.ToLower(e.Name), q) {
			out = append(out, e)
		}
	}
	return out
}

// ByDepartment returns every employee in dept.
func (m *Manager) ByDepartment(dept string) []Employee {
	d, _ := NormalizeDept(dept)
	var out []Employee
	for _, e := range m.list {
		if e.Department == d {
			out = append(out, e)
		}
	}
	return out
}

// Update changes the department and salary of an existing employee.
// An empty dept or a salary <= 0 leaves that field unchanged.
func (m *Manager) Update(id int, dept string, salary float64) (Employee, error) {
	pos, ok := m.index[id]
	if !ok {
		return Employee{}, fmt.Errorf("%w: ID %d", ErrNotFound, id)
	}
	if dept != "" {
		d, ok := NormalizeDept(dept)
		if !ok {
			return Employee{}, fmt.Errorf("%w: unknown department %q", ErrInvalid, dept)
		}
		m.list[pos].Department = d
	}
	if salary > 0 {
		m.list[pos].Salary = salary
	}
	return m.list[pos], nil
}

// Delete removes an employee from the slice and rebuilds the index,
// because every record after it has moved one position left.
func (m *Manager) Delete(id int) error {
	pos, ok := m.index[id]
	if !ok {
		return fmt.Errorf("%w: ID %d", ErrNotFound, id)
	}
	m.list = append(m.list[:pos], m.list[pos+1:]...)
	delete(m.index, id)
	for i := pos; i < len(m.list); i++ {
		m.index[m.list[i].ID] = i
	}
	return nil
}

// GiveRaise increases the salary of everyone in dept (or everyone when
// dept is "all") by percent and returns how many employees were updated.
func (m *Manager) GiveRaise(dept string, percent float64) (int, error) {
	if percent <= 0 || percent > 100 {
		return 0, fmt.Errorf("%w: percent must be between 0 and 100", ErrInvalid)
	}
	all := strings.EqualFold(dept, "all")
	d, ok := NormalizeDept(dept)
	if !all && !ok {
		return 0, fmt.Errorf("%w: unknown department %q", ErrInvalid, dept)
	}
	count := 0
	for i := range m.list { // index loop so we change the real element
		if all || m.list[i].Department == d {
			m.list[i].Salary *= 1 + percent/100
			count++
		}
	}
	return count, nil
}

// SortField selects how List orders employees.
type SortField int

// Sort options for List.
const (
	ByID SortField = iota
	ByName
	BySalaryDesc
	ByExperienceDesc
)

// List returns a sorted copy of all employees; the stored order is unchanged.
func (m *Manager) List(by SortField) []Employee {
	out := make([]Employee, len(m.list))
	copy(out, m.list)
	sort.SliceStable(out, func(i, j int) bool {
		switch by {
		case ByName:
			return out[i].Name < out[j].Name
		case BySalaryDesc:
			return out[i].Salary > out[j].Salary
		case ByExperienceDesc:
			return out[i].Experience > out[j].Experience
		default:
			return out[i].ID < out[j].ID
		}
	})
	return out
}

// Count returns the number of employees.
func (m *Manager) Count() int { return len(m.list) }

// Stats summarises salaries across all employees.
type Stats struct {
	Count                int
	Total, Avg, Min, Max float64
	Highest, Lowest      Employee
}

// Summary computes salary statistics with a single loop.
func (m *Manager) Summary() Stats {
	s := Stats{Count: len(m.list)}
	for i, e := range m.list {
		s.Total += e.Salary
		if i == 0 || e.Salary > s.Max {
			s.Max, s.Highest = e.Salary, e
		}
		if i == 0 || e.Salary < s.Min {
			s.Min, s.Lowest = e.Salary, e
		}
	}
	if s.Count > 0 {
		s.Avg = s.Total / float64(s.Count)
	}
	return s
}

// DeptStat is one row of the department report.
type DeptStat struct {
	Department string
	Count      int
	Total, Avg float64
}

// DepartmentReport groups employees by department using a map, then
// returns the rows sorted by department name.
func (m *Manager) DepartmentReport() []DeptStat {
	groups := make(map[string]*DeptStat)
	for _, e := range m.list {
		g, ok := groups[e.Department]
		if !ok {
			g = &DeptStat{Department: e.Department}
			groups[e.Department] = g
		}
		g.Count++
		g.Total += e.Salary
	}
	out := make([]DeptStat, 0, len(groups))
	for _, g := range groups {
		g.Avg = g.Total / float64(g.Count)
		out = append(out, *g)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Department < out[j].Department })
	return out
}

// NormalizeDept matches a department case-insensitively ("hr" -> "HR").
func NormalizeDept(dept string) (string, bool) {
	dept = strings.TrimSpace(dept)
	for _, d := range Departments {
		if strings.EqualFold(d, dept) {
			return d, true
		}
	}
	return "", false
}
