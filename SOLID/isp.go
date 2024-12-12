package solid

import "fmt"

// the payment interface contains refund in which some payment methods doesn't provide refund so the interface should be segregated
type Payment interface {
	Pay(amount int)
	Refund(amoutn int)
}

type Refund interface {
	Refund(amount int)
}

type PaymentType struct {
	paymentMethod PaymentMethod
}

func (p *PaymentType) Payment(amount int) {
	p.paymentMethod.Pay(amount)
}

type RefundService struct {
	refund Refund
}

func (p *RefundService) Refund(amount int) {
	p.refund.Refund(amount)
}

type CreditCard struct{}

func (p *CreditCard) Pay(amount int) {
	fmt.Printf("Amount %v paid using CreditCard \n", amount)
}

func (p *CreditCard) Refund(amount int) {
	fmt.Printf("Amount %v refunded using CreditCard \n", amount)
}

type DebitCard struct{}

func (p *DebitCard) Pay(amount int) {
	fmt.Printf("Amount %v paid using GooglePay \n", amount)
}

func isp() {
	creditCard := &CreditCard{}
	payment := &PaymentType{paymentMethod: creditCard}
	payment.paymentMethod.Pay(100)

	debitCard := &DebitCard{}
	dPayment := &PaymentType{paymentMethod: debitCard}
	dPayment.paymentMethod.Pay(100)

	refundService := &RefundService{refund: creditCard}
	refundService.refund.Refund(50)
}
