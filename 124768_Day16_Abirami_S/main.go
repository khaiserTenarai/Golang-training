package main

import (
	"log"
	"net/http"

	"day16-employee-graphql/config"
	"day16-employee-graphql/database"
	graphqlserver "day16-employee-graphql/graphql"
	"day16-employee-graphql/repository"
	"day16-employee-graphql/service"
)

func main() {
	cfg := config.Load()

	db := database.New()
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)

	server, err := graphqlserver.NewServer(employeeService)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", graphqlserver.NewHTTPHandler(server))

	log.Printf("GraphQL server running at http://localhost:%s/graphql", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
