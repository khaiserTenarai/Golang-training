package main

import "fmt"

func studentInfo() (string, int, float32) {
	name := "Tom"
	age := 15
	var marks float32 = 90
	return name, age, marks
}
func main() {
	name, age, marks := studentInfo()
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Marks: ", marks)
}
