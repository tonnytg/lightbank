package fixture

import (
	"context"
	"database/sql"
	"github.com/maragudk/migrate"
	"io/fs"
	"log"
)

func Up(migrationDir fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	if err := migrate.Up(context.Background(), db, migrationDir); err != nil {
		log.Fatal(err)
	}
	return db
}

func Down(db *sql.DB, migrationDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationDir); err != nil {
		log.Fatal(err)
	}
}
