package main
import "fmt"
var name ="Global"
func main(){
	fmt.Println("Global Variable: ",name)
	name2 := "Local"
	fmt.Println("Local Variable: ",name2)
	
}
