package main

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/factory-method/factory"
	"github.com/rafaelsanzio/go-patterns/factory-method/interfacefactory"
)

func main() {
	ak47, err := factory.GetGun("ak47")
	if err != nil {
		panic(err)
	}
	musket, err := factory.GetGun("musket")
	if err != nil {
		panic(err)
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g interfacefactory.IGun) {
	fmt.Println("Gun:", g.GetName())
	fmt.Println("Power:", g.GetPower())
}
