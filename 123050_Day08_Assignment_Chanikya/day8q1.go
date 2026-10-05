package main

import "fmt"

/*code besore fix
func divide(a int, b int) int {
	return a / b
}
panic: runtime error: integer divide by zero*/
//code after fixing the error
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
