package factory

import "fmt"

type Influx struct {
	url string
}

func (i *Influx) Connect() error {
	fmt.Println("Connecting to influx db", i.url)
	return nil
}
