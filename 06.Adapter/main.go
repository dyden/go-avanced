package main

import (
	. "github.com/dyden/go-avanced/06.Adapter/payment"
)

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	// bank := &BankPayment{}
	// ProcessPayment(bank)

	bpa := &BankPaymentAdapter{
		BankAccount: 1234,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
