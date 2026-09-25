package main

func main() {

	sum, diff, div := calc(5, 3)

	println("sum:", sum)
	println("diff:", diff)
	println("div:", div)
}

func calc(a int, b float64) (int, int, float64) {

	sum := a + int(b)
	diff := a - int(b)
	div := float64(a) / b

	return sum, diff, div
}
