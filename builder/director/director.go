package director

import (
	"github.com/rafaelsanzio/go-patterns/builder/interfacebuilder"
	"github.com/rafaelsanzio/go-patterns/builder/product"
)

type Director struct {
	builder interfacebuilder.IBuilder
}

func NewDirector(b interfacebuilder.IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b interfacebuilder.IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() product.House {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()

	return d.builder.GetHouse()
}
