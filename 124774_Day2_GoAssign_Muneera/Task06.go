package main
import "fmt"

func main(){
	var age int 
	fmt.Println("Enter your age")
	fmt.Scanf("%d",&age)

	if age>=18 && age<=60{
		fmt.Println("You are Eligible")
	}else{
		fmt.Println("You are not Eligible")
	}
}