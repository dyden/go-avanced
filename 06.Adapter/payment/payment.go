package payment

import "fmt"

type Payment interface {
	Pay()
}

//CASH
type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

//BANK
type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bankaccount: %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	BankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.BankAccount)
}
