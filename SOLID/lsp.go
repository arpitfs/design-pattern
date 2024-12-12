package solid

import "fmt"

type Gold struct{}

func (b *Gold) Pay(amount int) {
	fmt.Printf("Amount Paid using gold : %v \n", amount)
}

type Silver struct{}

func (u *Silver) Pay(amount int) {
	fmt.Printf("Amount Paid using silver : %v \n", amount)
}

type CommodityMethod interface {
	Pay(amount int)
}

type CommodityProcessor struct{}

func (p *CommodityProcessor) Payment(amount int, methodType Method) {
	methodType.Pay(amount)
}

func lsp() {
	gold := &Gold{}
	silver := &Silver{}
	processor := &CommodityProcessor{}

	processor.Payment(100, gold) // this can be replaced by silver or any new method
	processor.Payment(50, silver)
}
