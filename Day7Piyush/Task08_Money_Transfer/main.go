package main

import (
	"task08_money_transfer/config"
	"task08_money_transfer/controller"
	"task08_money_transfer/repository"
	"task08_money_transfer/view"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	repo := repository.NewTransferRepository(db)
	v := view.NewTransferView()
	ctrl := controller.NewTransferController(repo, v)

	ctrl.Run()
}
