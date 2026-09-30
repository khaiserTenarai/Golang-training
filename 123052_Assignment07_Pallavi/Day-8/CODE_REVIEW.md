var db map[string]string

func GetSalary(id string, base float64, bonus float64) float64 {
	x := base + base*bonus/100
	if id == "" {
		panic("no id")
	}
	return x
}

func main() {
	f, _ := os.Open("emp.txt")
	go func() { db["a"] = "b" }()
	fmt.Println(GetSalary("1", 1000, 10))
	fmt.Printf("%d", "x")
}