package main

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lighting signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
