package solid

import "fmt"

type Bank struct{}

func (b *Bank) Pay(amount int) {
	fmt.Printf("Amount Paid using bank : %v \n", amount)
}

type Upi struct{}

func (u *Upi) Pay(amount int) {
	fmt.Printf("Amount Paid using upi : %v \n", amount)
}

type Method interface {
	Pay(amount int)
}

type Processor struct{}

func (p *Processor) Payment(amount int, methodType Method) {
	methodType.Pay(amount)
}

// new methods can be added without modifying the existing code base
// create new payment method an implement method interface
func ocp() {
	bank := &Bank{}
	upi := &Upi{}
	processor := &Processor{}

	processor.Payment(100, bank)
	processor.Payment(50, upi)
}
