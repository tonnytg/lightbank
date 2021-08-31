package main

import (
	"github.com/tonnytg/lightbank/api"
	"github.com/tonnytg/lightbank/database"
)

func main() {

	// To prepare database with user and accounts, uncomment the next line
	// migrations.Migrate()

	// To prepare database for transactions, uncomment the next line
	// migrations.MigrateTransaction()

	// Start database Pool
	database.InitDatabase()

	// To start API Server
	api.StartApi()
}
