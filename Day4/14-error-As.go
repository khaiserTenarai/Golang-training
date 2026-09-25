package main 
import ( 
"errors" 
"fmt" 
) 
func main() { 
err := errors.New("employee not found") 
var target error 
if errors.As(err, &target) { 
fmt.Println("Error found:", target) 
} 
}