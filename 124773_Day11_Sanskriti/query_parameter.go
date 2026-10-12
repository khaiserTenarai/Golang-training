package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		name:= r.URL.Query().Get("name")
		page:= r.URL.Query().Get("page")

		fmt.Fprintln(w,"name",name,"page",page)
	})
	fmt.Println("connection success http://localhost:8080/employees")
	http.ListenAndServe(":8080", nil)
	
}