package repository
import "q12-repository-service-architecture/model"
type EmployeeRepository interface{CreateEmployee(model.Employee)error;GetEmployee(int)(*model.Employee,error);GetAllEmployees()([]model.Employee,error);UpdateEmployee(model.Employee)error;DeleteEmployee(int)error}
