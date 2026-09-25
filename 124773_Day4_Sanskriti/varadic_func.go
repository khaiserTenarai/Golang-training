package main
import "fmt"

func sal(salaries ...float64) float64{
	total := 0.0
	for _,salary := range salaries{
		total += salary
	}

	return total
}

func main(){
	total := sal(3000,5000,6000)
	fmt.Println("total",total)
}