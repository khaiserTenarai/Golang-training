package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"bankaccount/config"
	"bankaccount/controller"
	"bankaccount/repository"
	"bankaccount/service"
	"bankaccount/view"
)

func main() {

	fmt.Println("Starting Bank Account System...")

	cfg := config.Load()

	fmt.Println("Database:", cfg.DBName)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	db, err := pgxpool.New(
		ctx,
		cfg.DatabaseURL(),
	)

	if err != nil {
		log.Fatal(
			"Unable to create database pool:",
			err,
		)
	}

	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(
			"Unable to connect to PostgreSQL:",
			err,
		)
	}

	fmt.Println(
		"PostgreSQL connected successfully.",
	)

	// Repository

	accountRepository :=
		repository.NewAccountRepository(db)

	// Service

	accountService :=
		service.NewAccountService(
			accountRepository,
		)

	// View

	accountView :=
		view.NewAccountView()

	// Controller

	accountController :=
		controller.NewAccountController(
			accountService,
			accountView,
		)

	// Start application

	accountController.Start()
}
