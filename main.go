package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Retry is an exponential backoff retry helper. It is used to wait for postgres to boot up
func Retry(op func() error) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = time.Second * 10
	bo.MaxElapsedTime = time.Minute * 5

	if err := backoff.Retry(op, bo); err != nil {
		if bo.NextBackOff() == backoff.Stop {
			return fmt.Errorf("reached retry deadline")
		}

		return fmt.Errorf("retry failed: %w", err)
	}

	return nil
}

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

	var m *migrate.Migrate

	if err := Retry(func() error {
		var err error

		fmt.Println("trying migration...")

		m, err = migrate.New("file://migrations", uri)
		if err != nil {
			fmt.Printf("migration failed: (%+v) \n", err)
			return err
		}

		fmt.Println("migration completed successfully")

		return nil
	}); err != nil {
		log.Fatal("error creating new migration: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("error running migrations: ", err)
	}
}
