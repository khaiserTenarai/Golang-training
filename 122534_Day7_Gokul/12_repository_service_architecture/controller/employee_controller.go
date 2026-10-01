package controller

import (
	"bufio"
	"context"
	"example.com/q12-repository-service/model"
	"example.com/q12-repository-service/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EmployeeController struct {
	service service.EmployeeService
	reader  *bufio.Reader
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s, reader: bufio.NewReader(os.Stdin)}
}
func (c *EmployeeController) Start() {
	for {
		fmt.Println("\n1. Add Employee\n2. List Employees\n3. Exit")
		ch := c.i("Choice: ")
		ctx := context.Background()
		switch ch {
		case 1:
			e := model.Employee{Name: c.s("Name: "), Email: c.s("Email: "), Department: c.s("Department: ")}
			c.done(c.service.Add(ctx, e))
		case 2:
			list, err := c.service.List(ctx)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			for _, e := range list {
				fmt.Printf("%d | %s | %s | %s\n", e.ID, e.Name, e.Email, e.Department)
			}
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (c *EmployeeController) done(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Employee added.")
	}
}
func (c *EmployeeController) s(p string) string {
	fmt.Print(p)
	v, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(v)
}
func (c *EmployeeController) i(p string) int { v, _ := strconv.Atoi(c.s(p)); return v }
