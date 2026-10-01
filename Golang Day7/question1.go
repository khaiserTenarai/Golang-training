package main

import "fmt"

func main() {
   
    var fruitPrices map[string]float64

    
    fruitPrices["apple"] = 1.99

    fmt.Println("Apple price:", fruitPrices["apple"])
}

//panic: assignment to entry in nil map