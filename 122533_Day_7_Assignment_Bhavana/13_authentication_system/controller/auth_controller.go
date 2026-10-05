package controller

import (
	"bufio"
	"context"
	"example.com/q13-authentication/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AuthController struct {
	service service.AuthService
	reader  *bufio.Reader
}

func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{service: s, reader: bufio.NewReader(os.Stdin)}
}
func (c *AuthController) Start() {
	for {
		fmt.Println("\n1. Register\n2. Login\n3. Exit")
		ch := c.i("Choice: ")
		ctx := context.Background()
		switch ch {
		case 1:
			n := c.s("Username: ")
			p := c.s("Password: ")
			r := c.s("Role (user/admin): ")
			if e := c.service.Register(ctx, n, p, r); e != nil {
				fmt.Println("Error:", e)
			} else {
				fmt.Println("Registration successful.")
			}
		case 2:
			n := c.s("Username: ")
			p := c.s("Password: ")
			r, e := c.service.Login(ctx, n, p)
			if e != nil {
				fmt.Println("Login failed:", e)
			} else {
				fmt.Println("Login successful. Role:", r)
			}
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (c *AuthController) s(p string) string {
	fmt.Print(p)
	v, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(v)
}
func (c *AuthController) i(p string) int { v, _ := strconv.Atoi(c.s(p)); return v }
