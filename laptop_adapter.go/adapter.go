package main

import "fmt"

type adapter struct {
	windows *windows
}

func (a *adapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signal to USB signal.")
	a.windows.insertIntoUSBPort()
}
