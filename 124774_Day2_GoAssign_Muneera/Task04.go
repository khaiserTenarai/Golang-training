package main
import (
	"fmt"
	"strconv"
)
func main(){
	var a string="25"
	n,err := strconv.Atoi(a)
	if err != nil{
		fmt.Println("Invalid")
		return
	}
	b := float64(n)
	fmt.Println("String",a)
	fmt.Println("Integer",n)
	fmt.Println("Float",b)
}