package employee

import (
	"fmt"
	"github.com/fatih/color"
)

func EmpDetail(){

	fmt.Println("\n5. Go Module")
	name := "Vimal Aditya"
	age := 23
	id := 123041

	fmt.Println("Employee Name: ", name, "Age: ", age, "ID: ", id)

	fmt.Println("\n*******************************************")
	fmt.Println("\n6. External Package ")
	labelStyle := color.New(color.FgCyan, color.Bold)
	valStyle := color.New(color.FgGreen)

	fmt.Print("Employee Name: ")
	labelStyle.Print(name)

	fmt.Print(" | Age: ")
	valStyle.Print(age)

	fmt.Print(" | ID: ")
	valStyle.Println(id)

}