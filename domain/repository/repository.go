package repository

type TransactionRepository interface {
	Insert(id string, account string, amount string, status string, errorMessage string) error
}

