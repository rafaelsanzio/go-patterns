package factory

import "github.com/rafaelsanzio/go-patterns/abstract-factory/abstract"

// Concrete factory
type Adidas struct{}

func (a *Adidas) MakeShoe() abstract.IShoe {
	return &AdidasShoe{
		Shoe: abstract.Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() abstract.IShirt {
	return &AdidasShirt{
		Shirt: abstract.Shirt{
			Logo: "Adidas",
			Size: 14,
		},
	}
}
