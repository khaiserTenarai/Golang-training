package service

import (
 "context"
 "errors"
 "net/mail"
 "strings"

 "employee-management/model"
 "employee-management/repository"
)

var ErrInvalid = errors.New("invalid employee data")
type EmployeeService interface {
 Create(context.Context, model.EmployeeInput) (model.Employee, error)
 List(context.Context, int, int, string, string) ([]model.Employee, int64, error)
 Get(context.Context, int64) (model.Employee, error)
 Update(context.Context, int64, model.EmployeeInput) (model.Employee, error)
 Patch(context.Context, int64, model.EmployeePatch) (model.Employee, error)
 Delete(context.Context, int64) error
}
type employeeService struct { repo repository.EmployeeRepository }
func New(r repository.EmployeeRepository) EmployeeService { return &employeeService{repo:r} }

func validate(e model.Employee) error {
 if strings.TrimSpace(e.Name)=="" || len(e.Name)>100 || e.Age<18 || e.Salary<0 { return ErrInvalid }
 if _,err:=mail.ParseAddress(e.Email); err!=nil { return ErrInvalid }
 return nil
}
func toEmployee(in model.EmployeeInput) model.Employee { return model.Employee{Name:strings.TrimSpace(in.Name), Email:strings.TrimSpace(in.Email), Age:in.Age, Salary:in.Salary} }
func (s *employeeService) Create(ctx context.Context, in model.EmployeeInput) (model.Employee,error) {
 e:=toEmployee(in); if err:=validate(e); err!=nil{return e,err}; return s.repo.Create(ctx,e)
}
func (s *employeeService) List(ctx context.Context,p,l int,q,sort string)([]model.Employee,int64,error){ return s.repo.List(ctx,p,l,strings.TrimSpace(q),sort) }
func (s *employeeService) Get(ctx context.Context,id int64)(model.Employee,error){ return s.repo.Get(ctx,id) }
func (s *employeeService) Update(ctx context.Context,id int64,in model.EmployeeInput)(model.Employee,error){
 e:=toEmployee(in); if err:=validate(e);err!=nil{return e,err}; return s.repo.Update(ctx,id,e)
}
func (s *employeeService) Patch(ctx context.Context,id int64,p model.EmployeePatch)(model.Employee,error){
 e,err:=s.repo.Get(ctx,id); if err!=nil{return e,err}
 if p.Name!=nil { e.Name=strings.TrimSpace(*p.Name) }; if p.Email!=nil { e.Email=strings.TrimSpace(*p.Email) }; if p.Age!=nil { e.Age=*p.Age }; if p.Salary!=nil { e.Salary=*p.Salary }
 if p.Name==nil && p.Email==nil && p.Age==nil && p.Salary==nil { return e,ErrInvalid }
 if err:=validate(e);err!=nil{return e,err}
 return s.repo.Update(ctx,id,e)
}
func (s *employeeService) Delete(ctx context.Context,id int64)error{return s.repo.Delete(ctx,id)}
