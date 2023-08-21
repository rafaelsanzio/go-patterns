package main

import (
	"github.com/rafaelsanzio/go-patterns/structural/adapter/adapters"
	"github.com/rafaelsanzio/go-patterns/structural/adapter/client"
	"github.com/rafaelsanzio/go-patterns/structural/adapter/service"
)

func main() {
	cli := &client.Client{}
	mac := &service.Mac{}

	cli.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &service.Windows{}
	windowsMachineAdapter := &adapters.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	cli.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
