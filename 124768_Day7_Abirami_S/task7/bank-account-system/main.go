package main

import (
	"bank-account-system/controller"
	"bank-account-system/database"
	"bank-account-system/repository"
	"bank-account-system/service"
	"bank-account-system/view"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)
	accountView := view.NewAccountView(accountController)

	accountView.Start()
}
