package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("contert-type","text/plain")
	method := r.Method

	switch method{
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w,"Get request received")
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w,"Post request recived")

	case http.MethodPut:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w,"Put req received")
	
	case http.MethodDelete:
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w,"Method not allowed")
	}
	
}


func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server running at http://localhost:10004")

	err := http.ListenAndServe(":10004", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

