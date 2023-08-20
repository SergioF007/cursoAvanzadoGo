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

// Lo anterior es el primer compartamiento del metodo que esta implementado
// lo que hicimos es que mediante una interfaz implementamos  el metodo Pay del struc CashPayment

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
// como vemos esta funcion espera es el adaptador. que esta construido con su objeto principal
// y el valor que va a resivir el metodo que se pretende adaptar
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount) // aquie estariamos creando la sobreescritura del metodo
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
