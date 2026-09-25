package main 
 
import "fmt" 
 
func changeValue(x int) { 
 x = 100 
} 
 
func changePointer(x *int) { 
 *x = 100 
} 
 
func main() { 
 a := 10 
 b := 10 
 
 changeValue(a) 
 changePointer(&b) 
 
 fmt.Println("Value:", a) 
 fmt.Println("Pointer:", b) 
}