package main

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/abstract-factory/abstract"
	"github.com/rafaelsanzio/go-patterns/abstract-factory/interfacefactory"
)

func main() {
	adidasFactory, err := interfacefactory.GetSportsFactory("adidas")
	if err != nil {
		fmt.Println(err)
	}
	nikeFactory, err := interfacefactory.GetSportsFactory("nike")
	if err != nil {
		fmt.Println(err)
	}

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s abstract.IShoe) {
	fmt.Println("Logo:", s.GetLogo())
	fmt.Println("Size:", s.GetSize())
}

func printShirtDetails(s abstract.IShirt) {
	fmt.Println("Logo:", s.GetLogo())
	fmt.Println("Size:", s.GetSize())
}
