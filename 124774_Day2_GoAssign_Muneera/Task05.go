package main
import "fmt"
 func main(){
	var sal float64

	fmt.Println("Enter your salary")
	fmt.Scanf("%f",&sal)

	allowance :=sal*20/100
	bonus := sal*10/100
	deduction :=sal*5/100

	gross := sal+allowance+bonus
	net := gross-deduction

	fmt.Println("Basic Salary",sal)
	fmt.Println("Allowance",allowance)
	fmt.Println("Bonus",sal)
	fmt.Println("Deduction",deduction)
	fmt.Println("Gross Salary",gross)
	fmt.Println("Net Salary",net)
 }