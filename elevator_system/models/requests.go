package models

import "fmt"

type request struct {
	id        int
	fromFloor int
	toFloor   int
}

func CreateNewRequest(id int, from int, to int) *request {

	return &request{
		id:        id,
		fromFloor: from,
		toFloor:   to,
	}

}

func (req *request) GetRequestedDirection() (elevatorState, error) {

	requestedDirection := idle
	if req.fromFloor > req.toFloor {
		requestedDirection = down
	} else if req.fromFloor < req.toFloor {
		requestedDirection = up
	} else {
		return invalidState, fmt.Errorf("invalid floor")
	}

	return requestedDirection, nil
}
