package controller

import (
	"fmt"

	"assignment_task13/model"

	"assignment_task13/service"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(
	service service.AuthService,
) *AuthController {

	return &AuthController{

		service: service,
	}

}

func (c *AuthController) Register() {

	user := model.User{

		Username: "Vimal",

		Password: "12345",

		Role: "ADMIN",
	}

	err := c.service.Register(user)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println("Registration successful")

}

func (c *AuthController) Login() {

	err := c.service.Login(

		"Vimal",

		"12345",
	)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println("Login successful")

}
