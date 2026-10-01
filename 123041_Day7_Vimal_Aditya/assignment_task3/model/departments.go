package model

import "fmt"

type Department struct {
	ID   int64
	Name string
	Code string
}

func (d Department) Display() {
	fmt.Println("----------------------------------------")
	fmt.Println("Department ID   :", d.ID)
	fmt.Println("Department Name :", d.Name)
	fmt.Println("Department Code :", d.Code)
	fmt.Println("----------------------------------------")
}