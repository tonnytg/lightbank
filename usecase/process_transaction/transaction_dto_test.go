package process_transaction

import (
	"github.com/golang/mock/gomock"

	"github.com/tonnytg/lightbank/domain/entity/transactions"
	mock_repository "github.com/tonnytg/lightbank/domain/repository/mock"
	"testing"
	"time"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "1111111111111111",
		CreditCardName:            "Test Besta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	// manage mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)

	output, err := usecase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}

func TestProcessTransaction_ExecuteRejectTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4311111111111111",
		CreditCardName:            "Test Besta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.REJECTED,
		ErrorMessage: "you dont have limit for this transactions",
	}

	// manage mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)

	output, err := usecase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4311111111111111",
		CreditCardName:            "Test Besta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    900,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.APPROVED,
		ErrorMessage: "",
	}

	// manage mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)

	output, err := usecase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}