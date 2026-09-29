package main
import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println("fact", fact)

}

func factorial(n int) int{

	if n <= 1 {return 1}
	return n*factorial(n-1)

}