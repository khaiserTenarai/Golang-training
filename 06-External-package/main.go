package main

import "fmt"

import "github.com/google/uuid"
import "github.com/fatih/color"

func main(){

	id := uuid.New()
	fmt.Println(id)

}



// PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment> go get github.com/google/uuid
// go: added github.com/google/uuid v1.6.0





func main2() {
	fmt.Println("Student Application")

	color.Green("Student registration completed successfully!")
	color.Blue("Welcome to the Go application.")
}

/*
Before Using External packages , first we need to add them to our module
*/
/*
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy> go get  github.com/fatih/color        
go: downloading github.com/fatih/color v1.19.0
go: downloading github.com/mattn/go-isatty v0.0.20
go: downloading golang.org/x/sys v0.42.0
go: downloading github.com/mattn/go-colorable v0.1.14
go: added github.com/fatih/color v1.19.0
go: added github.com/mattn/go-colorable v0.1.14
go: added github.com/mattn/go-isatty v0.0.20
go: added golang.org/x/sys v0.42.0
*/


