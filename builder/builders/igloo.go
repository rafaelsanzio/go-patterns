package builders

import "github.com/rafaelsanzio/go-patterns/builder/product"

// Concrete igloo builder
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (ib *IglooBuilder) SetWindowType() {
	ib.windowType = "Snow Window"
}

func (ib *IglooBuilder) SetDoorType() {
	ib.doorType = "Snow Door"
}

func (ib *IglooBuilder) SetNumFloor() {
	ib.floor = 2
}

func (ib *IglooBuilder) GetHouse() product.House {
	return product.House{
		WindowType: ib.windowType,
		DoorType:   ib.doorType,
		Floor:      ib.floor,
	}
}
