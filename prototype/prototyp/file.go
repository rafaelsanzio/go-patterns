package prototyp

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/prototype/interfaceprototype"
)

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() interfaceprototype.INode {
	return &File{Name: f.Name + "_clone"}
}
