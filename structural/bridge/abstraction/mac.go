package abstraction

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/structural/bridge/implementation"
)

type Mac struct {
	printer implementation.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request from mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p implementation.Printer) {
	m.printer = p
}
