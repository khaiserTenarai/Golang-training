package service
import("q12-repository-service-architecture/model";"q12-repository-service-architecture/repository")
type EmployeeService interface{CreateEmployee(model.Employee)error;GetEmployee(int)(*model.Employee,error);GetAllEmployees()([]model.Employee,error);UpdateEmployee(model.Employee)error;DeleteEmployee(int)error}
type EmployeeServiceImpl struct{repository repository.EmployeeRepository}
func NewEmployeeService(r repository.EmployeeRepository)EmployeeService{return &EmployeeServiceImpl{r}}
func(s *EmployeeServiceImpl)CreateEmployee(e model.Employee)error{return s.repository.CreateEmployee(e)}
func(s *EmployeeServiceImpl)GetEmployee(id int)(*model.Employee,error){return s.repository.GetEmployee(id)}
func(s *EmployeeServiceImpl)GetAllEmployees()([]model.Employee,error){return s.repository.GetAllEmployees()}
func(s *EmployeeServiceImpl)UpdateEmployee(e model.Employee)error{return s.repository.UpdateEmployee(e)}
func(s *EmployeeServiceImpl)DeleteEmployee(id int)error{return s.repository.DeleteEmployee(id)}
