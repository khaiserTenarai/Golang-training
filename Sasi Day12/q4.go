package main

import "fmt"

type Notifier interface {
	Send(msg string)
}

type EmailService struct{}

func (e *EmailService) Send(msg string) {
	fmt.Println("Sending Email:", msg)
}

type UserHandler struct {
	notifier Notifier
}

func NewUserHandler(n Notifier) *UserHandler {
	return &UserHandler{notifier: n}
}

func (u *UserHandler) NotifyUser() {
	u.notifier.Send("Welcome to the system!")
}

func main() {
	emailSvc := &EmailService{}
	handler := NewUserHandler(emailSvc)
	handler.NotifyUser()
}