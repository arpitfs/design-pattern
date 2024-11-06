package factory

import "fmt"

type Postgres struct {
	url string
}

func (p *Postgres) Connect() error {
	fmt.Println("Connecting to postgres db", p.url)
	return nil
}
