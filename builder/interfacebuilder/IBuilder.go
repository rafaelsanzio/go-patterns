package interfacebuilder

import (
	"github.com/rafaelsanzio/go-patterns/builder/builders"
	"github.com/rafaelsanzio/go-patterns/builder/product"
)

// The builder interface
type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() product.House
}

func GetBuilder(bType string) IBuilder {
	switch bType {
	case "normal":
		return builders.NewNormalBuilder()
	case "igloo":
		return builders.NewIglooBuilder()
	default:
		return nil
	}
}
