package factory

import (
	"github.com/tonnytg/lightbank/domain/repository"
)

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
