package main

import "fmt"

type Employee struct {
    ID int
    Name string
}

type EmployeeRepository interface {
    Add(Employee)
    GetByID(int)
    Delete(int)
}

type Repository struct {
    employees []Employee
}

func (r *Repository) Add(emp Employee) {
    r.employees = append(r.employees, emp)
}

func (r *Repository) GetByID(id int) {
    for _, emp := range r.employees {
        if emp.ID == id {
            fmt.Println(emp.Name)
        }
    }
}

func (r *Repository) Delete(id int) {
    fmt.Println("Delete employee:", id)
}

func main() {

    var repo EmployeeRepository = &Repository{}

    repo.Add(Employee{ID: 101, Name: "Ganesh"})
    repo.GetByID(101)
    repo.Delete(101)
}


/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy\09-emp-interface-and-implementaion> go run .\Employee_Interface_Impl.go
Ganesh
Delete employee: 101
*/