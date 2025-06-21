package models

import "fmt"

type Building struct {
	AllFloors          []*floor
	elevatorController *elevatorController
}

func GetNewBuildingInstance(totalFloors int) *Building {

	elevatorController := getNewElevatorController(totalFloors)

	floorsList := make([]*floor, 0)

	for i := 0; i < totalFloors; i++ {

		floor := createFloor(i, createFloorDisplay(i, elevatorController, totalFloors))
		floorsList = append(floorsList, floor)

	}

	return &Building{
		AllFloors:          floorsList,
		elevatorController: elevatorController,
	}

}

func (i *Building) ReceiveRequestFromFloor(req *request) {

	i.elevatorController.receiveNewRequest(req)

}

func (i *Building) PrintPendingRequests() {
	fmt.Println(i.elevatorController.pendingRequests)
}
