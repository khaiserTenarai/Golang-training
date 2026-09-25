package main

import "fmt"

func main() {
	basic := 30000.0

	hra, da, pf := basic*0.40, basic*0.10, basic*0.12
	gross := basic + hra + da
	net := gross - pf

	fmt.Println("Basic:", basic)
	fmt.Println("HRA:", hra)
	fmt.Println("DA :", da)
	fmt.Println("PF :", pf)
	fmt.Println("Gross:", gross)
	fmt.Println("Net  :", net)
}