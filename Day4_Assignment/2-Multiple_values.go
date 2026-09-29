package main
import "fmt"

func main() {

name, age := employee("vittesh", 23)
fmt.Println("Name:", name)
fmt.Println("Age:", age)
}

func employee(name string, age int) (string, int) {
	return name, age

}