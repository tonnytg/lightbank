package process_transaction

import (
	"github.com/tonnytg/lightbank/domain/entity/credit_card"
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
	transaction := transactions.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	cc, invalidCC := creditcard.NewCreditCard(input.CreditCardNumber, input.CreditCardName,
		input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	// if invalidCC enter here
	if invalidCC != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, transactions.REJECTED, invalidCC.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}
		output := TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       transactions.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}
		return output, nil
	}
	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, transactions.REJECTED, invalidTransaction.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}
		output := TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       transactions.REJECTED,
			ErrorMessage: invalidTransaction.Error(),
		}
		return output, nil
	}

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, transactions.APPROVED, "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       transactions.APPROVED,
		ErrorMessage: "",
	}
	return output, nil
}
