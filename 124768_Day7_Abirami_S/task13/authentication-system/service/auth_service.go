package service
import("errors";"q13-authentication-system/model";"q13-authentication-system/repository";"golang.org/x/crypto/bcrypt")
type AuthService struct{repository repository.UserRepository}
func NewAuthService(r repository.UserRepository)*AuthService{return &AuthService{r}}
func(s *AuthService)Register(u,p,role string)error{if u==""||p==""{return errors.New("username and password are required")};h,e:=bcrypt.GenerateFromPassword([]byte(p),bcrypt.DefaultCost);if e!=nil{return e};return s.repository.CreateUser(model.User{Username:u,PasswordHash:string(h),Role:role})}
func(s *AuthService)Login(u,p string)(*model.User,error){x,e:=s.repository.GetByUsername(u);if e!=nil{return nil,e};if bcrypt.CompareHashAndPassword([]byte(x.PasswordHash),[]byte(p))!=nil{return nil,errors.New("invalid username or password")};return x,nil}
