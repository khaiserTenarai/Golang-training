package main

func main() {
    // Intentionally missing: import "fmt"
    fmt.Println("Hello, World!")
    /*
        PS C:\Users\mebona.joseph\Desktop\day1_go> go run builderror.go        
        builderror.go:1:1: expected 'package', found 'EOF'
        PS C:\Users\mebona.joseph\Desktop\day1_go> go build
        builderror.go:1:1: expected 'package', found 'EOF'

        Fix by importing "fmt"
    */

}