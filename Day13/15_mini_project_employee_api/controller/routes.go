package controller

import "net/http"

func RegisterRoutes(c *EmployeeController, apiKey string) http.Handler {
	auth := requireKey(apiKey)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", c.Health)
	mux.Handle("GET /employees", auth(http.HandlerFunc(c.GetAll)))
	mux.Handle("POST /employees", auth(http.HandlerFunc(c.Create)))
	mux.Handle("GET /employees/report", auth(http.HandlerFunc(c.Report)))
	mux.Handle("GET /employees/{id}", auth(http.HandlerFunc(c.GetOne)))
	mux.Handle("DELETE /employees/{id}", auth(http.HandlerFunc(c.Delete)))

	return mux
}
