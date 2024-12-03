package main

import (
	creational "designpattern/Creational"
	"designpattern/Creational/factory"
	solid "designpattern/SOLID"
	"fmt"
)

func main() {
	fmt.Println("Patterns")
	solid.Srp()
	factory.Factory()
	creational.BuildCreational()
}
