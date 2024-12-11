package facade

type PaymentMethod struct {
	paypal    *PaypalFacade
	googlePay *GooglePayFacade
	paytm     *PaytmFacade
}

func NewPaymentMethod() *PaymentMethod {
	return &PaymentMethod{
		paypal:    &PaypalFacade{},
		googlePay: &GooglePayFacade{},
		paytm:     &PaytmFacade{},
	}
}

func (p *PaymentMethod) PayWithPayPal(amount int) {
	p.paypal.Pay(amount)
}

func (p *PaymentMethod) PayWithGooglePay(amount int) {
	p.googlePay.Pay(amount)
}

func (p *PaymentMethod) PayWithPaytm(amount int) {
	p.paytm.Pay(amount)
}
