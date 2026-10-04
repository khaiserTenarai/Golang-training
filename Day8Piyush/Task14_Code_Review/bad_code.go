// Task 14: Code Review — This file has 10 deliberate issues.
// Review the code and find all 10 problems.
// See review_report.txt for the identified issues and fixes.
//
// NOTE: This file will NOT compile. It is meant for code review only.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Issue 1: Struct fields are unexported (lowercase) — can't be used from other packages
type employee struct {
	name   string
	salary float64
	age    int
}

// Issue 2: No input validation — negative salary and age are accepted
func createEmployee(name string, salary float64, age int) employee {
	return employee{name: name, salary: salary, age: age}
}

// Issue 3: Magic numbers — 0.30 and 10000 are unexplained constants
func calculateBonus(emp employee) float64 {
	if emp.salary > 50000 {
		return emp.salary * 0.30
	}
	return 10000
}

// Issue 4: Error is silently ignored — file close error and division error
func processSalaries(employees []employee) float64 {
	total := 0.0
	for _, emp := range employees {
		total += emp.salary
	}
	avg := total / float64(len(employees)) // Issue 5: Division by zero if slice is empty
	return avg
}

// Issue 6: Goroutine leak — goroutine writes to channel nobody reads
func fetchData() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "data" // blocks forever if nobody reads
	}()
	fmt.Println("fetchData returned without reading from channel")
}

// Issue 7: Race condition — shared map accessed by multiple goroutines without sync
func unsafeMapAccess() {
	data := make(map[int]string)
	for i := 0; i < 10; i++ {
		go func(val int) {
			data[val] = fmt.Sprintf("value_%d", val) // concurrent map write
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(data)
}

// Issue 8: Inefficient string concatenation in loop — should use strings.Builder
func buildReport(names []string) string {
	result := ""
	for _, name := range names {
		result += name + ", " // O(n²) string allocation
	}
	return result
}

// Issue 9: Mutex copied by value — receiver should be pointer
func (m sync.Mutex) badLock() {
	m.Lock()
	defer m.Unlock()
	fmt.Println("This copies the mutex, which is a bug")
}

// Issue 10: Error returned but never checked by caller
func getEmployeeByID(id int) (employee, error) {
	if id <= 0 {
		return employee{}, fmt.Errorf("invalid ID: %d", id)
	}
	return employee{name: "Piyush", salary: 50000, age: 25}, nil
}

func main() {
	emp, _ := getEmployeeByID(-1) // Issue 10: error ignored with _
	fmt.Println(emp)

	processSalaries([]employee{}) // Issue 5: empty slice causes division by zero

	fetchData() // Issue 6: goroutine leak

	unsafeMapAccess() // Issue 7: race condition
}
