package main

import (
	"github.com/rafaelsanzio/go-patterns/structural/composite/component"
	"github.com/rafaelsanzio/go-patterns/structural/composite/composition"
)

func main() {
	file1 := &component.File{Name: "File1"}
	file2 := &component.File{Name: "File2"}
	file3 := &component.File{Name: "File3"}

	folder1 := &composition.Folder{Name: "Folder1"}

	folder1.Add(file1)

	folder2 := &composition.Folder{Name: "Folder2"}

	folder2.Add(file2)
	folder2.Add(file3)

	folder2.Add(folder1)

	folder2.Search("rose")
}
