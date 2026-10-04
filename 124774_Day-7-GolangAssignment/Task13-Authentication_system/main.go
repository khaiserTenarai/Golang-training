package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"authentication/config"
	"authentication/controller"
	"authentication/repository"
	"authentication/service"
	"authentication/view"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	fmt.Println("Starting Authentication System...")

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

	userRepository :=
		repository.NewUserRepository(db)

	userService :=
		service.NewUserService(
			userRepository,
		)

	userView :=
		view.NewUserView()

	userController :=
		controller.NewUserController(
			userService,
			userView,
		)

	userController.Start()
}
