package main
import "fmt"
func main(){
	ch:=make(chan string, 3)
	ch<-"Employee 1"
	ch<-"Employee 2"
	ch<-"Employee 3"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}