package main
import "fmt"
func add(a int,b int){
	fmt.Println("Addition",a+b)
}

func sub(a int,b int){
	fmt.Println("Subtraction",a-b)
}

func mul(a int,b int){
	fmt.Println("Multiplication",a*b)
}

func div(a int,b int){
	fmt.Println("Division",a/b)
}

func main(){
	var num1 int
	var num2 int

	fmt.Println("Enter First Number")
	fmt.Scan(&num1)
	fmt.Println("Enter Second Number")
	fmt.Scan(&num2)

	add(num1,num2)
	sub(num1,num2)
	mul(num1,num2)
	div(num1,num2)
}