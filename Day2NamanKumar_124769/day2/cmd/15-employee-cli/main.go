// Command 15-employee-cli is the Day 2 mini project: an in-memory
// Employee Management CLI built with slices, maps, loops, conditions
// and functions.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"day2/internal/employee"
)

func main() {
	m := employee.NewSeeded()
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Employee Management CLI (5 sample employees loaded)")

	for {
		showMenu()
		switch ask(in, "Choose an option: ") {
		case "1":
			addEmployee(in, m)
		case "2":
			listEmployees(in, m)
		case "3":
			searchEmployee(in, m)
		case "4":
			updateEmployee(in, m)
		case "5":
			deleteEmployee(in, m)
		case "6":
			filterByDept(in, m)
		case "7":
			giveRaise(in, m)
		case "8":
			reports(m)
		case "0", "q", "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose 0-8.")
		}
	}
}

func showMenu() {
	fmt.Println(`
============ EMPLOYEE MANAGEMENT ============
 1. Add employee          5. Delete employee
 2. List employees        6. Filter by department
 3. Search (ID or name)   7. Give salary raise
 4. Update employee       8. Reports & statistics
 0. Exit
=============================================`)
}

func addEmployee(in *bufio.Scanner, m *employee.Manager) {
	name := ask(in, "Name: ")
	dept := ask(in, fmt.Sprintf("Department %v: ", employee.Departments))
	salary, ok := askFloat(in, "Monthly salary: ")
	if !ok {
		return
	}
	exp, ok := askInt(in, "Experience (years): ")
	if !ok {
		return
	}
	e, err := m.Add(name, dept, salary, exp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Added employee with ID %d\n", e.ID)
	printTable([]employee.Employee{e})
}

func listEmployees(in *bufio.Scanner, m *employee.Manager) {
	if m.Count() == 0 {
		fmt.Println("No employees yet.")
		return
	}
	var by employee.SortField
	switch ask(in, "Sort by: 1) ID  2) Name  3) Salary  4) Experience [1]: ") {
	case "2":
		by = employee.ByName
	case "3":
		by = employee.BySalaryDesc
	case "4":
		by = employee.ByExperienceDesc
	default:
		by = employee.ByID
	}
	printTable(m.List(by))
}

func searchEmployee(in *bufio.Scanner, m *employee.Manager) {
	q := ask(in, "Enter employee ID or part of a name: ")
	if id, err := strconv.Atoi(q); err == nil { // numeric -> map lookup
		e, err := m.Get(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printTable([]employee.Employee{e})
		return
	}
	results := m.SearchByName(q) // text -> loop over the slice
	if len(results) == 0 {
		fmt.Printf("No employee name contains %q\n", q)
		return
	}
	printTable(results)
}

func updateEmployee(in *bufio.Scanner, m *employee.Manager) {
	id, ok := askInt(in, "Employee ID to update: ")
	if !ok {
		return
	}
	cur, err := m.Get(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	printTable([]employee.Employee{cur})
	dept := ask(in, "New department (Enter to keep): ")
	var salary float64
	if s := ask(in, "New salary (Enter to keep): "); s != "" {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil || v <= 0 {
			fmt.Println("Error: salary must be a positive number")
			return
		}
		salary = v
	}
	e, err := m.Update(id, dept, salary)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Updated:")
	printTable([]employee.Employee{e})
}

func deleteEmployee(in *bufio.Scanner, m *employee.Manager) {
	id, ok := askInt(in, "Employee ID to delete: ")
	if !ok {
		return
	}
	e, err := m.Get(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if c := strings.ToLower(ask(in, fmt.Sprintf("Delete %s? (y/n): ", e.Name))); c != "y" && c != "yes" {
		fmt.Println("Cancelled.")
		return
	}
	if err := m.Delete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deleted %s (ID %d). %d employees left.\n", e.Name, id, m.Count())
}

func filterByDept(in *bufio.Scanner, m *employee.Manager) {
	dept := ask(in, fmt.Sprintf("Department %v: ", employee.Departments))
	if _, ok := employee.NormalizeDept(dept); !ok {
		fmt.Println("Error: unknown department")
		return
	}
	list := m.ByDepartment(dept)
	if len(list) == 0 {
		fmt.Println("No employees in that department.")
		return
	}
	printTable(list)
}

func giveRaise(in *bufio.Scanner, m *employee.Manager) {
	dept := ask(in, "Department (or 'all'): ")
	pct, ok := askFloat(in, "Raise percent: ")
	if !ok {
		return
	}
	n, err := m.GiveRaise(dept, pct)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Gave a %.1f%% raise to %d employee(s).\n", pct, n)
}

func reports(m *employee.Manager) {
	s := m.Summary()
	if s.Count == 0 {
		fmt.Println("No employees yet.")
		return
	}
	fmt.Println("\n------------- SALARY SUMMARY -------------")
	fmt.Printf("Employees      : %d\n", s.Count)
	fmt.Printf("Monthly payroll: ₹%s\n", money(s.Total))
	fmt.Printf("Average salary : ₹%s\n", money(s.Avg))
	fmt.Printf("Highest        : ₹%s (%s)\n", money(s.Max), s.Highest.Name)
	fmt.Printf("Lowest         : ₹%s (%s)\n", money(s.Min), s.Lowest.Name)

	fmt.Println("\n------------ DEPARTMENT REPORT ------------")
	fmt.Printf("%-12s %5s %14s %12s\n", "Department", "Count", "Total", "Average")
	for _, d := range m.DepartmentReport() {
		fmt.Printf("%-12s %5d %14s %12s\n", d.Department, d.Count, money(d.Total), money(d.Avg))
	}

	fmt.Println("\n------------ EXPERIENCE LEVELS ------------")
	levels := map[string]int{}
	for _, e := range m.List(employee.ByID) {
		levels[level(e.Experience)]++
	}
	for _, l := range []string{"Junior", "Mid", "Senior"} {
		fmt.Printf("%-7s %s %d\n", l, strings.Repeat("■", levels[l]), levels[l])
	}
}

// level classifies experience with conditions.
func level(years int) string {
	switch {
	case years < 3:
		return "Junior"
	case years < 6:
		return "Mid"
	default:
		return "Senior"
	}
}

func printTable(list []employee.Employee) {
	fmt.Printf("%-4s %-16s %-12s %12s %4s  %s\n", "ID", "Name", "Department", "Salary", "Exp", "Level")
	fmt.Println(strings.Repeat("-", 62))
	for _, e := range list {
		fmt.Printf("%-4d %-16s %-12s %12s %4d  %s\n",
			e.ID, e.Name, e.Department, money(e.Salary), e.Experience, level(e.Experience))
	}
	fmt.Printf("(%d record(s))\n", len(list))
}

// money formats a number with Indian digit grouping: 1234567 -> 12,34,567.
func money(v float64) string {
	s := strconv.FormatInt(int64(v+0.5), 10)
	if len(s) <= 3 {
		return s
	}
	head, tail := s[:len(s)-3], s[len(s)-3:]
	var parts []string
	for len(head) > 2 {
		parts = append([]string{head[len(head)-2:]}, parts...)
		head = head[:len(head)-2]
	}
	parts = append([]string{head}, parts...)
	return strings.Join(parts, ",") + "," + tail
}

func ask(in *bufio.Scanner, label string) string {
	fmt.Print(label)
	if !in.Scan() {
		fmt.Println("\nGoodbye!")
		os.Exit(0)
	}
	return strings.TrimSpace(in.Text())
}

func askInt(in *bufio.Scanner, label string) (int, bool) {
	v, err := strconv.Atoi(ask(in, label))
	if err != nil {
		fmt.Println("Error: please enter a whole number")
		return 0, false
	}
	return v, true
}

func askFloat(in *bufio.Scanner, label string) (float64, bool) {
	v, err := strconv.ParseFloat(ask(in, label), 64)
	if err != nil {
		fmt.Println("Error: please enter a number")
		return 0, false
	}
	return v, true
}
