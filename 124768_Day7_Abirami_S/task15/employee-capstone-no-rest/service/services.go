package service
import("errors";"q15-employee-capstone-no-rest/model";"q15-employee-capstone-no-rest/repository";"golang.org/x/crypto/bcrypt")
type AuthService struct{r *repository.UserRepository}
func NewAuthService(r *repository.UserRepository)*AuthService{return &AuthService{r}}
func(s *AuthService)Register(u,p,role string)error{h,e:=bcrypt.GenerateFromPassword([]byte(p),bcrypt.DefaultCost);if e!=nil{return e};return s.r.Create(model.User{Username:u,PasswordHash:string(h),Role:role})}
func(s *AuthService)Login(u,p string)(*model.User,error){x,e:=s.r.Get(u);if e!=nil{return nil,e};if bcrypt.CompareHashAndPassword([]byte(x.PasswordHash),[]byte(p))!=nil{return nil,errors.New("invalid credentials")};return x,nil}
type EmployeeService struct{r *repository.EmployeeRepository}
func NewEmployeeService(r *repository.EmployeeRepository)*EmployeeService{return &EmployeeService{r}}
func(s *EmployeeService)Create(e model.Employee)error{return s.r.Create(e)}
func(s *EmployeeService)Get(id int)(*model.Employee,error){return s.r.Get(id)}
func(s *EmployeeService)List(page,size int,name string)([]model.Employee,error){return s.r.List(page,size,name)}
func(s *EmployeeService)UpdateSalary(id int,x float64)error{return s.r.UpdateSalary(id,x)}
