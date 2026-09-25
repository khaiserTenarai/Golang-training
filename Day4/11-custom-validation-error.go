package main 
 
import "fmt" 
 
func validate(age int) error { 
 if age < 18 { 
  return fmt.Errorf("age must be 18 or above") 
 } 
 return nil 
} 
 
func main() { 
 err := validate(15) 
 
 if err != nil { 
  fmt.Println(err) 
 } 
} 