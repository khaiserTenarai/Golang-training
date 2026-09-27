package dao

import (
	"fmt"
	"employee_management/models"
)

type EmployeeDaoImpl struct{
	employees []models.Employee
}

func NewEmployeeDao() EmployeeDao {
	return &EmployeeDaoImpl{
		employees: []models.Employee{},
	}
}

func (r *EmployeeDaoImpl) Save(employee models.Employee) error{
	fmt.Println("Hello from DAO: Save Employee", employee.Name)

	for _, emp := range r.employees{
		if emp.ID == employee.ID{
			return fmt.Errorf("Employee with ID %d already exists", employee.ID)
		}
	}
	r.employees = append(r.employees, employee)
	return nil
}

func (r *EmployeeDaoImpl) FindById(id int) (models.Employee, error){
	fmt.Println("Hello from DAO: Find Employee ID", id)

	for _, emp := range r.employees{
		if emp.ID == id{
			return emp, nil
		}
	}
	return models.Employee{}, fmt.Errorf("Employee with ID %d not found", id)
}

func (r *EmployeeDaoImpl) FindAll() ([]models.Employee, error){
	fmt.Println("Hello from DAO: Find All Employee")

	return r.employees, nil
}

func (r *EmployeeDaoImpl) Update(employee models.Employee) error{
	fmt.Println("Hello from DAO: Update Employee", employee.Name)

	for i:=0; i < len(r.employees); i++{
		if r.employees[i].ID == employee.ID{
			r.employees[i] = employee
			return nil
		}
	}
	return fmt.Errorf("Employee with ID %d not found", employee.ID)
}

func (r *EmployeeDaoImpl) Delete(id int) error{
	fmt.Println("Hello from DAO: Delete from Emplpyee", id)

	for i := 0; i < len(r.employees); i++{
		if r.employees[i].ID == id{
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("employee with ID %d not found", id)
}

