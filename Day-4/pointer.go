package main

import "fmt"

func main() {
    x := 10
    p := &x       // p stores the address of x

    fmt.Println(x)  // 10
    fmt.Println(*p) // 10

    *p = 20        // changes x through the pointer
    fmt.Println(x) // 20
}