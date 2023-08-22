package abstraction

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/structural/bridge/implementation"
)

type Windows struct {
	printer implementation.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p implementation.Printer) {
	w.printer = p
}
