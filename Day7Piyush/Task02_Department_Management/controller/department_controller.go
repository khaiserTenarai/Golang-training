package controller

import (
	"fmt"
	"task02_department_management/repository"
	"task02_department_management/view"
)

type DepartmentController struct {
	deptRepo *repository.DepartmentRepository
	empRepo  *repository.EmployeeRepository
	view     *view.DepartmentView
}

func NewDepartmentController(dr *repository.DepartmentRepository, er *repository.EmployeeRepository, v *view.DepartmentView) *DepartmentController {
	return &DepartmentController{deptRepo: dr, empRepo: er, view: v}
}

func (c *DepartmentController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.CreateDepartment()
		case 2:
			c.ListDepartments()
		case 3:
			c.UpdateDepartment()
		case 4:
			c.DeleteDepartment()
		case 5:
			c.AddEmployee()
		case 6:
			c.ListAllEmployees()
		case 7:
			c.ListEmployeesByDept()
		case 8:
			c.AssignDepartment()
		case 9:
			c.DeleteEmployee()
		case 10:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *DepartmentController) CreateDepartment() {
	dept := c.view.GetDepartmentInput()
	id, err := c.deptRepo.Create(dept)
	if err != nil {
		c.view.ShowError("creating department", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Department created with ID: %d", id))
}

func (c *DepartmentController) ListDepartments() {
	depts, err := c.deptRepo.GetAll()
	if err != nil {
		c.view.ShowError("listing departments", err)
		return
	}
	c.view.ShowDepartments(depts)
}

func (c *DepartmentController) UpdateDepartment() {
	id := c.view.GetID("Department")
	existing, err := c.deptRepo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching department", err)
		return
	}
	updated := c.view.GetUpdateDeptInput(existing)
	if err := c.deptRepo.Update(updated); err != nil {
		c.view.ShowError("updating department", err)
		return
	}
	c.view.ShowSuccess("Department updated.")
}

func (c *DepartmentController) DeleteDepartment() {
	id := c.view.GetID("Department")
	if err := c.deptRepo.Delete(id); err != nil {
		c.view.ShowError("deleting department", err)
		return
	}
	c.view.ShowSuccess("Department deleted. Employees set to unassigned.")
}

func (c *DepartmentController) AddEmployee() {
	emp := c.view.GetEmployeeInput()
	id, err := c.empRepo.Create(emp)
	if err != nil {
		c.view.ShowError("adding employee", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Employee added with ID: %d", id))
}

func (c *DepartmentController) ListAllEmployees() {
	employees, err := c.empRepo.GetAllWithDepartment()
	if err != nil {
		c.view.ShowError("listing employees", err)
		return
	}
	c.view.ShowEmployees(employees)
}

func (c *DepartmentController) ListEmployeesByDept() {
	deptID := c.view.GetID("Department")
	employees, err := c.empRepo.GetByDepartment(deptID)
	if err != nil {
		c.view.ShowError("listing employees by department", err)
		return
	}
	c.view.ShowEmployees(employees)
}

func (c *DepartmentController) AssignDepartment() {
	empID := c.view.GetID("Employee")
	deptID := c.view.GetID("Department")
	if err := c.empRepo.AssignDepartment(empID, deptID); err != nil {
		c.view.ShowError("assigning department", err)
		return
	}
	c.view.ShowSuccess("Employee assigned to department.")
}

func (c *DepartmentController) DeleteEmployee() {
	id := c.view.GetID("Employee")
	if err := c.empRepo.Delete(id); err != nil {
		c.view.ShowError("deleting employee", err)
		return
	}
	c.view.ShowSuccess("Employee deleted.")
}
