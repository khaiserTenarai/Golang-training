package main

import "fmt"

func main() {

	var basicPay float64
	var workers int
	var totalBonus float64

	fmt.Print("Enter basic salary: ")
	fmt.Scan(&basicPay)

	hra := basicPay * 0.40
	da := basicPay * 0.10
	pf := basicPay * 0.05

	gross := basicPay + hra + da
	net := gross - pf

	fmt.Println(">>>> Salary Slip ")
	fmt.Println("Basic Pay:", basicPay)
	fmt.Println("HRA:", hra)
	fmt.Println("DA:", da)
	fmt.Println("PF Cut:", pf)
	fmt.Println("Gross Pay:", gross)
	fmt.Println("Net Pay:", net)

	fmt.Print("Enter number of workers: ")
	fmt.Scan(&workers)

	fmt.Print("Enter total bonus: ")
	fmt.Scan(&totalBonus)

	bonus := totalBonus / float64(workers)
	left := int(totalBonus) % workers

	fmt.Println(">>>> Bonus Info ")
	fmt.Println("Bonus per person:", bonus)
	fmt.Println("Leftover money:", left)
}
