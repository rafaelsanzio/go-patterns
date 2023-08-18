package adapters

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/structural/adapter/service"
)

type WindowsAdapter struct {
	WindowMachine *service.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
