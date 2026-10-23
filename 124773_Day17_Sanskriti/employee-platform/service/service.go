package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-platform/model"
	"employee-platform/repository"
	"employee-platform/utility"
)

type Service struct { Repo *repository.Repository }
func New(r *repository.Repository)*Service{return &Service{Repo:r}}
func (s *Service) Login(username,password string)(string,error){
	role,ok:=utility.DemoCredentials(username,password);if !ok{return "",fmt.Errorf("invalid credentials")}
	minutes:=60;if v:=os.Getenv("JWT_EXPIRES_MINUTES");v!=""{if n,e:=strconv.Atoi(v);e==nil && n>0{minutes=n}}
	return utility.GenerateToken(username,role,os.Getenv("JWT_SECRET"),minutes)
}
func (s *Service) ListEmployees(ctx context.Context,page,limit int,name,department,sortBy,order string)([]model.Employee,int,error){
	if page<1{return nil,0,fmt.Errorf("page must be at least 1")};if limit<1||limit>100{return nil,0,fmt.Errorf("limit must be between 1 and 100")}
	sortBy=strings.ToLower(sortBy);if sortBy==""{sortBy="id"}
	if sortBy!="id"&&sortBy!="name"&&sortBy!="email"&&sortBy!="age"&&sortBy!="salary"{return nil,0,fmt.Errorf("unsupported sort field")}
	order=strings.ToUpper(order);if order==""{order="ASC"};if order!="ASC"&&order!="DESC"{return nil,0,fmt.Errorf("order must be ASC or DESC")}
	return s.Repo.ListEmployees(ctx,page,limit,name,department,sortBy,order)
}
func(s *Service)GetEmployee(ctx context.Context,id int64)(model.Employee,error){return s.Repo.GetEmployee(ctx,id)}
func(s *Service)CreateEmployee(ctx context.Context,in model.EmployeeInput)(model.Employee,error){if err:=utility.ValidateEmployee(in);err!=nil{return model.Employee{},err};return s.Repo.CreateEmployee(ctx,in)}
func(s *Service)UpdateEmployee(ctx context.Context,id int64,in model.EmployeeInput)(model.Employee,error){if err:=utility.ValidateEmployee(in);err!=nil{return model.Employee{},err};return s.Repo.UpdateEmployee(ctx,id,in)}
func(s *Service)DeleteEmployee(ctx context.Context,id int64)error{return s.Repo.DeleteEmployee(ctx,id)}
func(s *Service)ListDepartments(ctx context.Context)([]model.Department,error){return s.Repo.ListDepartments(ctx)}
func(s *Service)GetDepartment(ctx context.Context,id int64)(model.Department,error){return s.Repo.GetDepartment(ctx,id)}
