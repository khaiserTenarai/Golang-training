package controller
import("q13-authentication-system/model";"q13-authentication-system/service")
type AuthController interface{Register(string,string,string)error;Login(string,string)(*model.User,error)}
type AuthControllerImpl struct{service *service.AuthService}
func NewAuthController(s *service.AuthService)AuthController{return &AuthControllerImpl{s}}
func(c *AuthControllerImpl)Register(a,b,d string)error{return c.service.Register(a,b,d)}
func(c *AuthControllerImpl)Login(a,b string)(*model.User,error){return c.service.Login(a,b)}
