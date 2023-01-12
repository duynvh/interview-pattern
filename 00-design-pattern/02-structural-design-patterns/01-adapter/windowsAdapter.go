package main

import "fmt"

type WindowAdapters struct {
	windowMachine *Windows
}

func (w *WindowAdapters) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
