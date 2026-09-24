package main

import (
	
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	
	fakeName := gofakeit.Name()
	fakeEmail := gofakeit.Email()
	fakeJob := gofakeit.JobTitle()

	
	fmt.Println("Name: ", fakeName)
	fmt.Println("Email:", fakeEmail)
	fmt.Println("Job:  ", fakeJob)
}