package main

import (
	creational "designpattern/Creational"
	"designpattern/Creational/factory"
	solid "designpattern/SOLID"
	structural "designpattern/Structural"
	"fmt"
)

func main() {
	fmt.Println("Patterns")
	solid.SOLID()
	factory.Factory()
	creational.BuildCreational()
	structural.Structural()
}
