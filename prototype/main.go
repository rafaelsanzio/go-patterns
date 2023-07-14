package main

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/prototype/interfaceprototype"
	"github.com/rafaelsanzio/go-patterns/prototype/prototyp"
)

func main() {
	file1 := &prototyp.File{Name: "File1"}
	file2 := &prototyp.File{Name: "File2"}
	file3 := &prototyp.File{Name: "File3"}

	folder1 := &prototyp.Folder{
		Children: []interfaceprototype.INode{file1},
		Name:     "Folder 1",
	}

	folder2 := &prototyp.Folder{
		Children: []interfaceprototype.INode{folder1, file2, file3},
		Name:     "Folder 2",
	}

	fmt.Println("Printing hierarchy for Folder 2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()

	fmt.Println("Printing hierarchy for clone Folder")
	cloneFolder.Print("  ")

}
