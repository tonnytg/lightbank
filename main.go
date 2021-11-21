package main

import (
	"github.com/tonnytg/lightbank/api"
	"github.com/tonnytg/lightbank/migrations"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		// Create Database
		migrations.Migrate()
	} else {
		// api migration
		api.StartApi()
	}
}
