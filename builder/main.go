package main

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/builder/director"
	"github.com/rafaelsanzio/go-patterns/builder/interfacebuilder"
)

func main() {
	normalBuilder := interfacebuilder.GetBuilder("normal")
	iglooBuilder := interfacebuilder.GetBuilder("igloo")

	direc := director.NewDirector(normalBuilder)
	normalHouse := direc.BuildHouse()

	fmt.Println("Normal House Door Type:", normalHouse.DoorType)
	fmt.Println("Normal House Window Type:", normalHouse.WindowType)
	fmt.Println("Normal House Num Floor:", normalHouse.Floor)

	direc.SetBuilder(iglooBuilder)
	iglooHouse := direc.BuildHouse()

	fmt.Println("Igloo House Door Type:", iglooHouse.DoorType)
	fmt.Println("Igloo House Window Type:", iglooHouse.WindowType)
	fmt.Println("Igloo House Num Floor:", iglooHouse.Floor)
}
