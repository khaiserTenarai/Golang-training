package main
import "fmt"

func greetWithDetails(name string, age int) {
	fmt.Println(name, age)
}

// Keeping the original greet name here
func greet() {
	greetWithDetails("Pallavi", 22)
}
func calculate(a,b int)(int,int){

	sum:=a+b
	difference :=a-b
	return sum,difference
}
//using fmt and without declaring datatype
func calculate1(a,b int){

	sum:=a+b
	difference :=a-b
	fmt.Println(sum,difference)
}

//using different datatype

func strore(a string ,b float64){
	
	fmt.Println(a,b)
	
	
}
func multiply(a,b int){
	value:=a*b
	diff:=a/b
	fmt.Println("value:",value)
	fmt.Println("diff:",diff)
}

//adding with range concepts
func add(numbers ...int)int{
	total:=0
	for _,number:=range numbers{
		total+=number
	}
	return total
}

// Every executable Go program needs a main function to run
func main() {
	greet()
	s, d := calculate(10, 5)
	multiply(10,5)
	// Print the values
	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
	calculate1(10,5)
	fmt.Println("dummy","Data")
	strore("Studio",12.5)
	multiply(10,5)
	fmt.Println(add(10,20))
	fmt.Println(add(10,30,40))
	func(){
		fmt.Println("I am anonymous!")
	}()
}


//anonymous function : without function name
