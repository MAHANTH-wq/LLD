package models

import (
	"fmt"
	"time"
)

type doorStatus int

const (
	open doorStatus = iota
	close
)

type doors struct {
	status doorStatus
}

func createDoors() *doors {
	return &doors{
		status: close,
	}
}
func (d *doors) getDoorStatus() doorStatus {

	return d.status

}
func (d *doors) openDoors() {
	fmt.Println("Opening Doors")
	d.status = open
	time.Sleep(2 * time.Second)
}

func (d *doors) closeDoors() {
	fmt.Println("Closing Doors")
	d.status = close
	time.Sleep(2 * time.Second)

}
