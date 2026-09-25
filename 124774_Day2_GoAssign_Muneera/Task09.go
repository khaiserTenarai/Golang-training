package main
import "fmt"
func main(){
	sales :=[12]float64{
		10000,12000,14000,16000,18000,20000,
		11000,13000,15000,17000,19000,21000,
	}
	total:=0.0
	for i:=0;i<12;i++{
		total+=sales[i]
	}
	average :=total/12
	fmt.Println("Total Sales:",total)
	fmt.Println("Average Sales:",average)
}