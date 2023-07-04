/* Say, you need to buy a sports kit, a set of two different products: a pair of shoes and a shirt. You would want to buy a full sports kit of the same brand to match all the items.

If we try to turn this into code, the abstract factory will help us create sets of products so that they would always match each other. */

package interfacefactory

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/abstract-factory/abstract"
	"github.com/rafaelsanzio/go-patterns/abstract-factory/factory"
)

// Abstract factory interface
type ISportsFactory interface {
	MakeShoe() abstract.IShoe
	MakeShirt() abstract.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &factory.Adidas{}, nil
	case "nike":
		return &factory.Nike{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}
