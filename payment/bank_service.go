package payment

import (
	"fmt"
	"golang-di-playground/policyholder"
)

type BankService struct{}

func NewBankService() Service {
	return &BankService{}
}

func (s *BankService) Pay(policyHolder *policyholder.PolicyHolder, amount float64) error {
	// Here, you would implement the logic to make a payment.
	// This could involve interacting with a bank API, a payment gateway, etc.
	// For this example, let's just print a message.
	fmt.Printf("Paying $%.2f by bank transfer, to policyholder %s\n", amount, policyHolder)
	return nil
}
