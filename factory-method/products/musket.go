package products

import "github.com/rafaelsanzio/go-patterns/factory-method/interfacefactory"

type Musket struct {
	Gun
}

func NewMusket() interfacefactory.IGun {
	return &Musket{
		Gun: Gun{
			Name:  "Musket",
			Power: 1,
		},
	}
}
