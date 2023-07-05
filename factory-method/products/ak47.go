package products

import "github.com/rafaelsanzio/go-patterns/factory-method/interfacefactory"

type Ak47 struct {
	Gun
}

func NewAk47() interfacefactory.IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "Ak47",
			Power: 4,
		},
	}
}
