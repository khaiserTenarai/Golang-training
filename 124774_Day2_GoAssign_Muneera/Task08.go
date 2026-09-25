package main
import "fmt"
func main(){
	name := "नमस्ते"
	bytes :=[]byte(name)

	fmt.Println("Name: ",name)
	fmt.Println("Bytes: ",bytes)
	fmt.Println("Byte Count: ",len(bytes))

	runes :=[]rune(name)

	fmt.Println("Runes:",runes)
	fmt.Println("Rune count",len(runes))
}