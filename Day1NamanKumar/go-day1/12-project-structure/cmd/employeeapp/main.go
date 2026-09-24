// Command employeeapp is the executable entry point. It only wires
// together config, storage and validation — no business logic lives here.
package main

import (
	"fmt"
	"log"

	"github.com/example/employeeapp/internal/config"
	"github.com/example/employeeapp/internal/store"
	"github.com/example/employeeapp/pkg/validator"
)

func main() {
	cfg := config.Load()
	fmt.Printf("%s starting (env=%s)\n", cfg.AppName, cfg.Env)

	s := store.New()
	email := "asha@corp.in"
	if err := validator.Email(email); err != nil {
		log.Fatal(err)
	}
	e := s.Add("Asha Rao", email)
	fmt.Printf("Added employee #%d %s <%s>\n", e.ID, e.Name, e.Email)
}
