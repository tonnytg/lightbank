package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/tonnytg/lightbank/adapter/repository/fixture"
	"github.com/tonnytg/lightbank/domain/entity/transactions"
	"os"
	"testing"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)
	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "2", 12.1, transactions.APPROVED, "")
	if err != nil {
		t.Error(err)
	}
}
