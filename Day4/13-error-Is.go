package main 
import ( 
"errors" 
"fmt" 
) 
var ErrNotFound = errors.New("employee not found") 
func getEmployee() error { 
return ErrNotFound 
} 
func main() { 
err := getEmployee() 
if errors.Is(err, ErrNotFound) { 
fmt.Println("Employee not found") 
} 
} 
