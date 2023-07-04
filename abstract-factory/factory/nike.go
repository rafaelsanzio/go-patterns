package factory

import "github.com/rafaelsanzio/go-patterns/abstract-factory/abstract"

// Concrete factory
type Nike struct{}

func (n *Nike) MakeShoe() abstract.IShoe {
	return &NikeShoe{
		Shoe: abstract.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() abstract.IShirt {
	return &NikeShirt{
		Shirt: abstract.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
