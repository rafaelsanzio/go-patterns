package builders

import "github.com/rafaelsanzio/go-patterns/builder/product"

// Concrete normal builder
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) SetWindowType() {
	nb.windowType = "Wooden Window"
}

func (nb *NormalBuilder) SetDoorType() {
	nb.doorType = "Wooden Door"
}

func (nb *NormalBuilder) SetNumFloor() {
	nb.floor = 2
}

func (nb *NormalBuilder) GetHouse() product.House {
	return product.House{
		WindowType: nb.windowType,
		DoorType:   nb.doorType,
		Floor:      nb.floor,
	}
}
