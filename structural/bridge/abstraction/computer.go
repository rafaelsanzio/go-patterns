package abstraction

import "github.com/rafaelsanzio/go-patterns/structural/bridge/implementation"

type Computer interface {
	Print()
	SetPrinter(implementation.Printer)
}
