package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightingConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightingConnectorIntoComputer(windowsMachineAdapter)
}
