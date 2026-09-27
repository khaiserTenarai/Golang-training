package main


type EmployeeRepository interface {
    Add(Employee)
    GetByID(int)
    Delete(int)
}