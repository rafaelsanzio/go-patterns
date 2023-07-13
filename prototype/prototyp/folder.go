package prototyp

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/prototype/interfaceprototype"
)

type Folder struct {
	Children []interfaceprototype.INode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)

	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() interfaceprototype.INode {
	var tempChildren []interfaceprototype.INode

	cloneFolder := &Folder{Name: f.Name + "_clone"}

	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.Children = tempChildren
	return cloneFolder
}
