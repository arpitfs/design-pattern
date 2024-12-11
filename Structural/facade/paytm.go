package facade

import "fmt"

type PaytmFacade struct{}

func (p *PaytmFacade) Pay(amount int) {
	fmt.Println("Amount Payed using Paytm: ", amount)
}
