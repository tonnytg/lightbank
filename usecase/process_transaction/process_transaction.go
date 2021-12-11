package usecase

import (
	"github.com/tonnytg/lightbank/domain/entity/transactions"
	"github.com/tonnytg/lightbank/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := transactions.NewTransaction(input.From, input.To, input.Amount)
	return
}