package main

import (
	"task07_bank_account_system/config"
	"task07_bank_account_system/controller"
	"task07_bank_account_system/repository"
	"task07_bank_account_system/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	repo := repository.NewAccountRepository(db)
	v := view.NewAccountView()
	ctrl := controller.NewAccountController(repo, v)

	ctrl.Run()
}
