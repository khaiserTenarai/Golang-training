// Task 5: Run gofmt — This file is INTENTIONALLY unformatted.
// Run: gofmt -d unformatted.go        (shows diff)
// Run: gofmt -w unformatted.go        (formats in place)

package main

import (
"fmt"
  "math"
    "strings"
)

func main(){
var     x    =10
  var y=     20
  result:=x+     y
fmt.Println(    "Sum:",result)

    if x>5{
fmt.Println("x is greater than 5")
  }else{
      fmt.Println("x is not greater")
}

names := []string{   "Alice","Bob",
"Charlie",   "David",
}
fmt.Println(strings.Join(      names,", "))

  area:=math.Pi*     5*5
    fmt.Printf("Circle area: %.2f\n",area)
}
