package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")

	if host == "" {
		log.Fatal("DATABASE_HOST not specified")
	}

	if port == "" {
		log.Fatal("DATABASE_PORT not specified")
	}

	if username == "" {
		log.Fatal("DATABASE_USERNAME not specified")
	}

	if password == "" {
		log.Fatal("DATABASE_PASSWORD not specified")
	}

	if name == "" {
		log.Fatal("DATABASE_NAME not specified")
	}

	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, name)

	m, err := migrate.New("file://migrations", uri)
	if err != nil {
		log.Fatal("error creating new migration: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("error running migrations: ", err)
	}
}
