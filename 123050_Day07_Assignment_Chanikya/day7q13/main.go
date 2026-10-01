package main

import (
	"context"

	"fmt"

	"github.com/jackc/pgx/v5"

	"day7q13/repository"

	"day7q13/service"

	"day7q13/controller"
)

func main() {

	conn, err := pgx.Connect(

		context.Background(),

		"postgres://postgres:admin@localhost:5432/gotraining",
	)

	if err != nil {

		panic(err)

	}

	defer conn.Close(context.Background())

	// Repository

	userRepository := repository.NewUserRepository(conn)

	// Service

	authService := service.NewAuthService(
		userRepository,
	)

	// Controller

	authController := controller.NewAuthController(
		authService,
	)

	fmt.Println("----- REGISTER -----")

	authController.Register()

	fmt.Println("\n----- LOGIN -----")

	authController.Login()

}
