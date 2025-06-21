package models

import (
	"fmt"
	"time"
)

type elevatorState int

const (
	up elevatorState = iota
	idle
	down
	invalidState
)

type elevatorVehicle struct {
	capacity     int
	maxCapacity  int
	state        elevatorState
	currentFloor int
	id           int
	display      display
	doors        *doors
}

func createElevator(id int, maxCapacity int) *elevatorVehicle {

	ev := &elevatorVehicle{
		capacity:     0,
		maxCapacity:  maxCapacity,
		state:        idle,
		currentFloor: 0,
		id:           id,
		doors:        createDoors(),
	}
	return ev
}

func (ev *elevatorVehicle) move(direction elevatorState, floorId int) {

	if direction == up {
		fmt.Println("Lift Moving Up from floor ", ev.currentFloor, " to floor ", floorId)
	} else if direction == down {
		fmt.Println("Lift Moving Down from floor ", ev.currentFloor, " to floor ", floorId)
	} else {
		fmt.Println("Lift is in Idle State")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Lift Reached floor ", floorId)
	ev.doors.openDoors()
	ev.doors.closeDoors()

}

func (ev *elevatorVehicle) checkDoors() doorStatus {
	return ev.doors.status
}

func (ev *elevatorVehicle) closeDoors() {
	ev.doors.closeDoors()
}

func (ev *elevatorVehicle) openDoors() {
	ev.doors.openDoors()
}
