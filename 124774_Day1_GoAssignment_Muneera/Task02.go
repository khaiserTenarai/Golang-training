package main
import "fmt"
func firstway(){
	fmt.Println("Application Name : Hello Cloud-Native Go" )
	fmt.Println("Application version : 1.0.0")
	fmt.Println("Go Version : go1.27.1")
	fmt.Println("Environment : Development")
	
}
func secondway(){
	application := "Hello Cloud-Native Go"
	version := "1.0.0"
	go_version := "go1.27.1"
	environment := "Development"

	fmt.Println("Application Name : " ,application)
	fmt.Println("Application version : " ,version)
	fmt.Println("Go Version : ",go_version)
	fmt.Println("Environment : ",environment)
}

func main(){
	firstway()
	secondway()
}