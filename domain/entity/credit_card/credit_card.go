package entity

import (
	"errors"
	"regexp"
)

type CreditCard struct {
	Number      string
	Cvv         int64
	ExpiryMonth int64
	ExpiryYear  int64
}

func NewCreditCard(code, name string, month, year, cvv int64) (CreditCard, error) {
	cc := CreditCard{}
	cc.Number = code
	cc.ExpiryMonth = month
	cc.ExpiryYear = year
	cc.Cvv = cvv

	err := cc.CheckNumber()
	if err != nil {
		return cc, err
	}

	err = cc.CheckCvv()
	if err != nil {
		return cc, err
	}

	err = cc.CheckExpiry()
	if err != nil {
		return cc, err
	}

	return cc, nil
}

func (c *CreditCard) CheckNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.Number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (c *CreditCard) CheckCvv() error {

	if c.Cvv < 100 || c.Cvv > 999 {
		return errors.New("invalid cvv")
	}
	return nil
}

func (c *CreditCard) CheckExpiry() error {
	// TODO: check year and month
	if c.ExpiryMonth < 1 || c.ExpiryMonth > 12 {
		return errors.New("invalid expiry month")
	}
	if c.ExpiryYear < 2021 || c.ExpiryYear > 2030 {
		return errors.New("invalid expiry year")
	}
	return nil
}
