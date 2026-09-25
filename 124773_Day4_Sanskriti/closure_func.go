package main

import "fmt"

func clo() func() int{
	id :=0

	return func() int{
		id++
		return id
	}

}

func main(){
	nextid := clo()

	fmt.Println(nextid())
}