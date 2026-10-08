package repository
import "q08-money-transfer/model"
type AccountRepository interface{GetAccount(int)(*model.Account,error);Transfer(int,int,float64)error}
