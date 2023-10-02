package composition

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/structural/composite/leaf"
)

type Folder struct {
	Components []leaf.Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)

	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c leaf.Component) {
	f.Components = append(f.Components, c)
}
