package main

import (
	"github.com/tonnytg/lightbank/api"
	"github.com/tonnytg/lightbank/migrations"
	"os"
)

func main() {

	// To prepare database with user and accounts, uncomment the next line
	// migrations.Migrate()

	// To prepare database for transactions, uncomment the next line
	// migrations.MigrateTransaction()
	if len(os.Args) > 1 {
		// Create Database
		migrations.Migrate()
	} else {
		// api migration
		api.StartApi()
	}
}
