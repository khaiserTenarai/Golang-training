package main

import (
    "fmt"
    "os"//"os" imported and not used
)

func main() {
    age := "25" 

    
    nextYear := age + 1 //mismatched types string and untyped int
    
    fmt.Println("Next year you will be:", nextYear)
}
/*package main

import (
    "fmt"
    
)

func main() {
    
    age := 25 

    nextYear := age + 1 
    
    fmt.Println("Next year you will be:", nextYear)
}*/