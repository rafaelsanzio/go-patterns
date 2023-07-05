package factory

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/factory-method/interfacefactory"
	"github.com/rafaelsanzio/go-patterns/factory-method/products"
)

func GetGun(gType string) (interfacefactory.IGun, error) {
	switch gType {
	case "ak47":
		return products.NewAk47(), nil
	case "musket":
		return products.NewMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
