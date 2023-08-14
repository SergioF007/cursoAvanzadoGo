package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d\n", bankAccount)
}

// Vamos a crear el adaptador del metodo Pay
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

// Vamos hacer que BankPaymentAdapter implemete el Pay de la manera correcta
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {

	cash := &CashPayment{}
	ProcessPayment(cash)
	// probamos

	//bank := &BankPayment{}
	//ProcessPayment(bank)

	// Instanciamos y creamos el adaptador .
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}

	ProcessPayment(bpa)

}
