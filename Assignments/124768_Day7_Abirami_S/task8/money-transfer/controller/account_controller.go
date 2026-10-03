package controller
import("q08-money-transfer/model";"q08-money-transfer/service")
type AccountController interface{GetAccount(int)(*model.Account,error);Transfer(int,int,float64)error}
type AccountControllerImpl struct{service service.AccountService}
func NewAccountController(s service.AccountService)AccountController{return &AccountControllerImpl{s}}
func(c *AccountControllerImpl)GetAccount(id int)(*model.Account,error){return c.service.GetAccount(id)}
func(c *AccountControllerImpl)Transfer(a,b int,x float64)error{return c.service.Transfer(a,b,x)}
