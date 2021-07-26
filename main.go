package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://trackmyfish:supersecretpassword@localhost:15432/trackmyfish?sslmode=disable",
	)
	if err != nil {
		log.Fatal("error creating new migration: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("error running migrations: ", err)
	}
}
