package entity_test

import (
	"github.com/tonnytg/lightbank/domain/entity/credit_card"
	"testing"
)

func TestCreditCardNumber(t *testing.T) {

	// Test to pass
	_, err := entity.NewCreditCard("5431111111111111", "Beta Tester", 12, 2024, 123)
	if err != nil {
		t.Errorf("Check bussiness rule ValidNumber. Error creating credit card: %s", err)
	}

	// Test to fail
	_, err = entity.NewCreditCard("1111111111111111", "Beta Tester", 12, 2024, 123)
	if err == nil {
		t.Errorf("Check bussiness rule ValidNumber. Invalid credit card passed: %s", err)
	}
}

func TestCreditCardExpirationMonth(t *testing.T) {

	// Test to pass
	_, err := entity.NewCreditCard("5431111111111111", "Beta Tester", 12, 2024, 123)
	if err != nil {
		t.Errorf("Check bussiness rule ValidMonth. Error creating credit card: %s", err)
	}

	// Test to fail
	_, err = entity.NewCreditCard("5431111111111111", "Beta Tester", 14, 2024, 123)
	if err == nil {
		t.Errorf("Check bussiness rule ValidMonth. Error creating credit card: %s", err)
	}
}

func TestCreditCardExpirationYear(t *testing.T) {

	// Test to pass
	_, err := entity.NewCreditCard("5431111111111111", "Beta Tester", 12, 2024, 123)
	if err != nil {
		t.Errorf("Check bussiness rule ValidYear. Error creating credit card: %s", err)
	}

	// Test to fail
	_, err = entity.NewCreditCard("5431111111111111", "Beta Tester", 12, 2020, 123)
	if err == nil {
		t.Errorf("Check bussiness rule ValidYear. Error creating credit card: %s", err)
	}
}
