package main 
import ( 
"errors" 
"fmt" 
) 
func main() { 
err := errors.New("file not found") 
newErr := fmt.Errorf("failed to open file: %w", err) 
fmt.Println(newErr) 
}