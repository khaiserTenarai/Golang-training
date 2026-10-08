package service
import("q08-money-transfer/model";"q08-money-transfer/repository")
type AccountService interface{GetAccount(int)(*model.Account,error);Transfer(int,int,float64)error}
type AccountServiceImpl struct{repository repository.AccountRepository}
func NewAccountService(r repository.AccountRepository)AccountService{return &AccountServiceImpl{r}}
func(s *AccountServiceImpl)GetAccount(id int)(*model.Account,error){return s.repository.GetAccount(id)}
func(s *AccountServiceImpl)Transfer(a,b int,x float64)error{return s.repository.Transfer(a,b,x)}
