package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Defaults
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.username", "trackmyfish")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.name", "trackmyfish")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("unable to read config")
	}

	var (
		dbHost     = viper.GetString("db.host")
		dbPort     = viper.GetInt("db.port")
		dbUsername = viper.GetString("db.username")
		dbPassword = viper.GetString("db.password")
		dbName     = viper.GetString("db.name")
	)

	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbPort, dbName)

	m, err := migrate.New("file://migrations", uri)
	if err != nil {
		log.Fatal("error creating new migration: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("error running migrations: ", err)
	}
}
