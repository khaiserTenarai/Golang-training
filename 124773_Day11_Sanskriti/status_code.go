package main

import "net/http" 

func main() {

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Success
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("Employees found"))
	})
}
	fmt.Println("http://localhost:8080/employee")
	http.ListenAndServe(":8080", nil)
}