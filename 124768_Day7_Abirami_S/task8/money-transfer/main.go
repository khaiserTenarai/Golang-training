package main
import("q08-money-transfer/controller";"q08-money-transfer/database";"q08-money-transfer/repository";"q08-money-transfer/service";"q08-money-transfer/view")
func main(){db:=database.ConnectDB();defer db.Close();r:=repository.NewAccountRepository(db);s:=service.NewAccountService(r);c:=controller.NewAccountController(s);view.NewAccountView(c).Start()}
