package service

import (
 "context"
 "secure-gin-api/model"
 "secure-gin-api/repository"
 "secure-gin-api/utility"
 "golang.org/x/crypto/bcrypt"
)
type EmployeeStore interface{Create(context.Context,model.EmployeeInput)(model.Employee,error);Get(context.Context,int64)(model.Employee,error);List(context.Context,int,int,string,string)([]model.Employee,int,error);Update(context.Context,int64,model.EmployeeInput)(model.Employee,error);Delete(context.Context,int64)error;FindUser(context.Context,string)(repository.User,error)}
type EmployeeService struct{Repo EmployeeStore}
func NewEmployeeService(r EmployeeStore)*EmployeeService{return &EmployeeService{Repo:r}}
func(s *EmployeeService)Create(ctx context.Context,e model.EmployeeInput)(model.Employee,error){if err:=utility.ValidateEmployee(e);err!=nil{return model.Employee{},err};return s.Repo.Create(ctx,e)}
func(s *EmployeeService)Get(ctx context.Context,id int64)(model.Employee,error){return s.Repo.Get(ctx,id)}
func(s *EmployeeService)List(ctx context.Context,p,l int,sort,order string)([]model.Employee,int,error){return s.Repo.List(ctx,p,l,sort,order)}
func(s *EmployeeService)Update(ctx context.Context,id int64,e model.EmployeeInput)(model.Employee,error){if err:=utility.ValidateEmployee(e);err!=nil{return model.Employee{},err};return s.Repo.Update(ctx,id,e)}
func(s *EmployeeService)Delete(ctx context.Context,id int64)error{return s.Repo.Delete(ctx,id)}
func(s *EmployeeService)Login(ctx context.Context,username,password string)(repository.User,error){u,err:=s.Repo.FindUser(ctx,username);if err!=nil{return repository.User{},err};if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash),[]byte(password))!=nil{return repository.User{},bcrypt.ErrMismatchedHashAndPassword};return u,nil}
