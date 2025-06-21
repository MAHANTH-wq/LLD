package models

import (
	utils "elevator_system/utils"
	"fmt"
	"sync"
)

type elevatorController struct {
	elevator                 *elevatorVehicle
	totalFloors              int
	pendingRequestsMutex     sync.Mutex
	handleConcurrentRequests sync.Mutex
	pendingRequests          []*request
	minHeapRequests          utils.Heap
	maxHeapRequests          utils.Heap
}

func getNewElevatorController(totalFloors int) *elevatorController {

	return &elevatorController{
		elevator:                 createElevator(0, 100),
		totalFloors:              totalFloors,
		pendingRequestsMutex:     sync.Mutex{},
		handleConcurrentRequests: sync.Mutex{},
		pendingRequests:          make([]*request, 0),
		minHeapRequests:          utils.CreateHeap(utils.MinHeapType),
		maxHeapRequests:          utils.CreateHeap(utils.MaxHeapType),
	}

}

func (ec *elevatorController) receiveNewRequest(req *request) error {

	fromFloor := req.fromFloor
	toFloor := req.toFloor
	requestedDirection, err := req.GetRequestedDirection()

	if err != nil {
		return err
	}

	flag := false

	ec.handleConcurrentRequests.Lock()

	if ec.elevator.state == idle {

		flag = true
		if requestedDirection == up {
			ec.elevator.state = up
			ec.minHeapRequests.InsertElement(fromFloor)
			ec.minHeapRequests.InsertElement(toFloor)
		} else {
			ec.elevator.state = down
			ec.maxHeapRequests.InsertElement(toFloor)
			ec.maxHeapRequests.InsertElement(fromFloor)
		}
	}

	ec.handleConcurrentRequests.Unlock()

	if flag == true {
		ec.processRequests()
		fmt.Println(ec.minHeapRequests.GetSize())
		fmt.Println(ec.maxHeapRequests.GetSize())
	} else {
		ec.pendingRequestsMutex.Lock()
		ec.pendingRequests = append(ec.pendingRequests, req)
		ec.pendingRequestsMutex.Unlock()
	}

	return nil
}

func (ec *elevatorController) processRequests() {

	for ec.minHeapRequests.GetSize() > 0 || ec.maxHeapRequests.GetSize() > 0 || len(ec.pendingRequests) > 0 {
		if ec.elevator.state == idle {

			if ec.minHeapRequests.GetSize() > 0 && ec.maxHeapRequests.GetSize() > 0 {

				if ec.elevator.currentFloor < (ec.totalFloors / 2) {
					ec.moveDown()
				} else {
					ec.moveUp()
				}

			} else if ec.minHeapRequests.GetSize() > 0 {
				ec.moveUp()

			} else if ec.maxHeapRequests.GetSize() > 0 {
				ec.moveDown()
			}
		} else if ec.elevator.state == down {
			ec.moveDown()
		} else if ec.elevator.state == up {
			ec.moveUp()
		}
	}

	ec.elevator.state = idle

}

func (ec *elevatorController) moveDown() {

	if ec.maxHeapRequests.GetSize() > 0 {
		if ec.elevator.currentFloor < ec.maxHeapRequests.GetTopElement() {
			ec.elevator.state = up
			ec.elevator.move(ec.elevator.state, ec.maxHeapRequests.GetTopElement())
			ec.elevator.currentFloor = ec.maxHeapRequests.GetTopElement()
			ec.maxHeapRequests.PopElement()
		}
		ec.elevator.state = down

		for ec.maxHeapRequests.GetSize() > 0 {
			topRequest := ec.maxHeapRequests.GetTopElement()
			ec.elevator.move(ec.elevator.state, topRequest)
			ec.elevator.currentFloor = topRequest
			ec.maxHeapRequests.PopElement()
		}

	}

	tempList := make([]*request, 0)

	ec.pendingRequestsMutex.Lock()
	for _, req := range ec.pendingRequests {

		if req.fromFloor <= req.toFloor {
			ec.minHeapRequests.InsertElement(req.toFloor)
			ec.minHeapRequests.InsertElement(req.fromFloor)
		} else {
			tempList = append(tempList, req)
		}

	}

	ec.pendingRequests = tempList

	ec.pendingRequestsMutex.Unlock()
	ec.elevator.state = up

}

func (ec *elevatorController) moveUp() {
	if ec.minHeapRequests.GetSize() > 0 {

		if ec.elevator.currentFloor > ec.minHeapRequests.GetTopElement() {
			ec.elevator.state = down
			ec.elevator.move(ec.elevator.state, ec.minHeapRequests.GetTopElement())
			ec.elevator.currentFloor = ec.minHeapRequests.GetTopElement()
			ec.minHeapRequests.PopElement()
		}
		ec.elevator.state = up

		for ec.minHeapRequests.GetSize() > 0 {
			topRequest := ec.minHeapRequests.GetTopElement()
			ec.elevator.move(ec.elevator.state, topRequest)
			ec.elevator.currentFloor = topRequest
			ec.minHeapRequests.PopElement()
		}

	}

	tempList := make([]*request, 0)
	ec.pendingRequestsMutex.Lock()

	for _, req := range ec.pendingRequests {

		if req.fromFloor >= req.toFloor {

			ec.maxHeapRequests.InsertElement(req.toFloor)
			ec.maxHeapRequests.InsertElement(req.fromFloor)
		} else {
			tempList = append(tempList, req)
		}

	}
	ec.pendingRequests = tempList
	ec.pendingRequestsMutex.Unlock()
	ec.elevator.state = down

}
