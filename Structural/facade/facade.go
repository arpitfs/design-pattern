package facade

func Facade() {
	paymentMethods := NewPaymentMethod()
	paymentMethods.PayWithPayPal(100)
	paymentMethods.PayWithGooglePay(200)
	paymentMethods.PayWithPaytm(300)
}
