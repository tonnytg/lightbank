package main

import (
	"github.com/tonnytg/lightbank/api"
	"github.com/tonnytg/lightbank/database"
	"github.com/tonnytg/lightbank/migrations"
	"os"
)

func init() {

	dbconf := database.DBConfg{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	database.InitDatabase(&dbconf)
}

func main() {

	if len(os.Args) > 1 {
		// Create Database
		migrations.Migrate()
	} else {
		// api migration
		api.StartApi()
	}
}
