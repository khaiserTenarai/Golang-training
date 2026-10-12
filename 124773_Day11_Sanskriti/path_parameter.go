package main

import (
	"fmt"
	"net/http"
	"strings"
)


func main() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		id:= strings.TrimPrefix(r.URL.Path,"/employees/")

		fmt.Fprintln(w,"Employee ID",id)
	})
	fmt.Println("connection success http://localhost:8080/employees")
	http.ListenAndServe(":8080", nil)
	
}