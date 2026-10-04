package controller

import (
	"fmt"

	"authentication/service"
	"authentication/view"
)

type UserController struct {
	service service.UserService
	view    *view.UserView
}

func NewUserController(
	service service.UserService,
	view *view.UserView,
) *UserController {

	return &UserController{
		service: service,
		view:    view,
	}
}

func (c *UserController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadString(
			"Enter your choice: ",
		)

		switch choice {

		case "1":
			c.Register()

		case "2":
			c.Login()

		case "3":
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
func (c *UserController) Register() {
	username := c.view.ReadString("Enter username: ")
	password := c.view.ReadString("Enter password: ")
	role := c.view.ReadString("Enter role (admin/user): ")
	err := c.service.Register(username, password, role)
	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error())
		return
	}
	c.view.ShowMessage("User registered successfully.")
}

func (c *UserController) Login() {

	username := c.view.ReadString(
		"Enter username: ",
	)

	password := c.view.ReadString(
		"Enter password: ",
	)

	err := c.service.Login(
		username,
		password,
	)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Login successful.",
	)
}
