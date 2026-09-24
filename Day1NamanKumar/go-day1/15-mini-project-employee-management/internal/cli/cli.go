// Package cli implements the interactive, menu-driven user interface.
// It reads from an io.Reader and writes to an io.Writer so it can be tested
// without a real terminal.
package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"

	"employeeapp/internal/employee"
)

// App is the interactive employee management application.
type App struct {
	store *employee.Store
	in    *bufio.Scanner
	out   io.Writer
}

// New creates an App using the given store and I/O streams.
func New(store *employee.Store, in io.Reader, out io.Writer) *App {
	return &App{store: store, in: bufio.NewScanner(in), out: out}
}

const menu = `
====== Employee Management ======
 1. Add employee
 2. Search employee
 3. Display all employees
 4. Delete employee
 5. Exit
=================================`

// Run shows the menu in a loop until the user exits or input ends.
func (a *App) Run() {
	for {
		fmt.Fprintln(a.out, menu)
		choice, ok := a.prompt("Choose an option (1-5): ")
		if !ok {
			return // EOF (e.g. Ctrl+D)
		}
		switch choice {
		case "1":
			a.add()
		case "2":
			a.search()
		case "3":
			a.display(a.store.List())
		case "4":
			a.delete()
		case "5", "q", "exit":
			fmt.Fprintln(a.out, "Goodbye!")
			return
		default:
			fmt.Fprintf(a.out, "Invalid option %q. Please enter 1-5.\n", choice)
		}
	}
}

// prompt prints label and returns the trimmed input line.
func (a *App) prompt(label string) (string, bool) {
	fmt.Fprint(a.out, label)
	if !a.in.Scan() {
		return "", false
	}
	return strings.TrimSpace(a.in.Text()), true
}

func (a *App) promptID() (int, bool) {
	s, ok := a.prompt("Employee ID: ")
	if !ok {
		return 0, false
	}
	id, err := strconv.Atoi(s)
	if err != nil || id <= 0 {
		fmt.Fprintf(a.out, "Error: %q is not a valid ID.\n", s)
		return 0, false
	}
	return id, true
}

func (a *App) add() {
	name, _ := a.prompt("Name: ")
	dept, _ := a.prompt("Department: ")
	salaryStr, _ := a.prompt("Salary: ")
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		fmt.Fprintf(a.out, "Error: %q is not a valid salary.\n", salaryStr)
		return
	}
	e, err := a.store.Add(employee.Employee{Name: name, Department: dept, Salary: salary})
	if err != nil {
		fmt.Fprintln(a.out, "Error:", err)
		return
	}
	fmt.Fprintf(a.out, "✔ Employee added with ID %d.\n", e.ID)
}

func (a *App) search() {
	mode, _ := a.prompt("Search by (1) ID or (2) Name: ")
	switch mode {
	case "1":
		id, ok := a.promptID()
		if !ok {
			return
		}
		e, err := a.store.GetByID(id)
		if errors.Is(err, employee.ErrNotFound) {
			fmt.Fprintf(a.out, "No employee found with ID %d.\n", id)
			return
		}
		a.display([]employee.Employee{e})
	case "2":
		q, _ := a.prompt("Name contains: ")
		a.display(a.store.SearchByName(q))
	default:
		fmt.Fprintln(a.out, "Invalid search option.")
	}
}

func (a *App) delete() {
	id, ok := a.promptID()
	if !ok {
		return
	}
	if err := a.store.Delete(id); err != nil {
		fmt.Fprintf(a.out, "Error: %v (ID %d).\n", err, id)
		return
	}
	fmt.Fprintf(a.out, "✔ Employee %d deleted.\n", id)
}

func (a *App) display(list []employee.Employee) {
	if len(list) == 0 {
		fmt.Fprintln(a.out, "No employees to display.")
		return
	}
	tw := tabwriter.NewWriter(a.out, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tNAME\tDEPARTMENT\tSALARY")
	fmt.Fprintln(tw, "--\t----\t----------\t------")
	for _, e := range list {
		fmt.Fprintf(tw, "%d\t%s\t%s\t%.2f\n", e.ID, e.Name, e.Department, e.Salary)
	}
	tw.Flush()
	fmt.Fprintf(a.out, "(%d record(s))\n", len(list))
}
