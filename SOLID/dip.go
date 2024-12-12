package solid

import "fmt"

type PaymentMethod interface {
	Pay(amount int)
}

type PayPal struct{}

func (p *PayPal) Pay(amount int) {
	fmt.Printf("Amount %v paid using Paypal \n", amount)
}

type GoogolePay struct{}

func (p *GoogolePay) Pay(amount int) {
	fmt.Printf("Amount %v paid using GooglePay \n", amount)
}

type Paytm struct{}

func (p *Paytm) Pay(amount int) {
	fmt.Printf("Amount %v paid using Paytm \n", amount)
}

type PaymentService struct {
	paymentMethod PaymentMethod
}

func (p *PaymentService) Payment(amount int) {
	p.paymentMethod.Pay(amount)
}

func dip() {
	paytm := &Paytm{}
	paytmPaymentService := &PaymentService{paymentMethod: paytm}
	paytmPaymentService.Payment(10)

	paypal := &PayPal{}
	paypalPaymentService := &PaymentService{paymentMethod: paypal}
	paypalPaymentService.Payment(10)
}
