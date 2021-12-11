package usecase

import (
	"github.com/golang/mock/gomock"
	entity "github.com/tonnytg/lightbank/domain/entity/credit_card"
	"github.com/tonnytg/lightbank/domain/entity/transactions"
	"github.com/golang/mock/mockgen"
	mock_repository "github.com/tonnytg/lightbank/domain/repository/mock"
	"testing"
	"time"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumer:           "5431111111111111",
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

	output, err := usercase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}