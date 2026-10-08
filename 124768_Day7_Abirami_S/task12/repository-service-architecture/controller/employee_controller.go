package controller
import("q12-repository-service-architecture/model";"q12-repository-service-architecture/service")
type EmployeeController interface{CreateEmployee(model.Employee)error;GetEmployee(int)(*model.Employee,error);GetAllEmployees()([]model.Employee,error);UpdateEmployee(model.Employee)error;DeleteEmployee(int)error}
type EmployeeControllerImpl struct{service service.EmployeeService}
func NewEmployeeController(s service.EmployeeService)EmployeeController{return &EmployeeControllerImpl{s}}
func(c *EmployeeControllerImpl)CreateEmployee(e model.Employee)error{return c.service.CreateEmployee(e)}
func(c *EmployeeControllerImpl)GetEmployee(id int)(*model.Employee,error){return c.service.GetEmployee(id)}
func(c *EmployeeControllerImpl)GetAllEmployees()([]model.Employee,error){return c.service.GetAllEmployees()}
func(c *EmployeeControllerImpl)UpdateEmployee(e model.Employee)error{return c.service.UpdateEmployee(e)}
func(c *EmployeeControllerImpl)DeleteEmployee(id int)error{return c.service.DeleteEmployee(id)}
