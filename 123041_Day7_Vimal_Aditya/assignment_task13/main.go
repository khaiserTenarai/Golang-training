package main

import (
	"context"

	"fmt"

	"github.com/jackc/pgx/v5"

	"assignment_task13/repository"

	"assignment_task13/service"

	"assignment_task13/controller"
)

func main() {

	conn, err := pgx.Connect(

		context.Background(),

		"postgres://postgres:Info%40131@localhost:5432/go_training",
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
