package utility

import (
	"fmt"
	"os"
	"strings"
	"time"

	"employee-platform/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(username, role, secret string, minutes int) (string,error) {
	claims:=jwt.MapClaims{"username":username,"role":role,"exp":time.Now().Add(time.Duration(minutes)*time.Minute).Unix(),"iat":time.Now().Unix()}
	t:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return t.SignedString([]byte(secret))
}
func ParseToken(raw,secret string)(*model.User,error){
	t,err:=jwt.Parse(raw,func(t *jwt.Token)(any,error){if t.Method != jwt.SigningMethodHS256{return nil,fmt.Errorf("unexpected signing method")};return []byte(secret),nil})
	if err!=nil || !t.Valid{return nil,fmt.Errorf("invalid token")}
	c,ok:=t.Claims.(jwt.MapClaims);if !ok{return nil,fmt.Errorf("invalid claims")}
	username,_:=c["username"].(string);role,_:=c["role"].(string)
	if username=="" || role=="" {return nil,fmt.Errorf("invalid claims")}
	return &model.User{Username:username,Role:role},nil
}
func ValidateEmployee(e model.EmployeeInput) error {
	if strings.TrimSpace(e.Name)=="" || len(e.Name)>100{return fmt.Errorf("name is required and must be at most 100 characters")}
	if _,err:=mailParse(e.Email);err!=nil{return fmt.Errorf("valid email is required")}
	if e.Age<18 || e.Age>100{return fmt.Errorf("age must be between 18 and 100")}
	if e.Salary<0{return fmt.Errorf("salary cannot be negative")}
	if e.DepartmentID<1{return fmt.Errorf("department_id must be positive")}
	return nil
}
func mailParse(s string)(string,error){if !strings.Contains(s,"@") || strings.HasPrefix(s,"@") || strings.HasSuffix(s,"@"){return "",fmt.Errorf("invalid email")};return s,nil}
func HashPassword(password string)(string,error){b,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost);return string(b),err}
func CheckPassword(hash,password string)bool{return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))==nil}
func DemoCredentials(username,password string)(string,bool){
	// Demo only. Configure real users and password hashes in production.
	u:=os.Getenv("ADMIN_USERNAME"); p:=os.Getenv("ADMIN_PASSWORD")
	if u=="" {u="admin"}; if p=="" {p="ChangeMe123!"}
	if username==u && password==p{return "admin",true}
	u=os.Getenv("USER_USERNAME"); p=os.Getenv("USER_PASSWORD")
	if u=="" {u="user"}; if p=="" {p="ChangeMe123!"}
	if username==u && password==p{return "user",true}
	return "",false
}
