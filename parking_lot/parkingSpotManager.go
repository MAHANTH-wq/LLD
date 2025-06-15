package main

import "fmt"

type parkingSpotManager interface {
	findParkingSpot() (parkingSpot, bool)
	parkVehicle(vehicle) (*ticket, bool)
	removeVehicle(*ticket)
	addParkingSpot(parkingSpot)
	removeParkingSpot(parkingSpot)
	getVehicleTypeForParkingSpotManager() string
}

func getNewParkingSpotManager(vType string) parkingSpotManager {

	if vType == "Two Wheeler" {
		return &twoWheelerParkingSpotManager{}
	} else if vType == "Four Wheeler" {
		return &fourWheelerParkingSpotManager{}
	} else {
		return nil
	}

}

type fourWheelerParkingSpotManager struct {
	listOfParkingSpots []parkingSpot
}

func (twm *fourWheelerParkingSpotManager) getVehicleTypeForParkingSpotManager() string {
	return "Four Wheeler"
}

func (twm *fourWheelerParkingSpotManager) addParkingSpot(ps parkingSpot) {
	twm.listOfParkingSpots = append(twm.listOfParkingSpots, ps)
}

func (twm *fourWheelerParkingSpotManager) removeParkingSpot(ps parkingSpot) {
	for index, lps := range twm.listOfParkingSpots {

		if ps.getParkingSpotId() == lps.getParkingSpotId() {
			twm.listOfParkingSpots = append(twm.listOfParkingSpots[0:index], twm.listOfParkingSpots[index+1:]...)
		}
	}
}
func (twm *fourWheelerParkingSpotManager) findParkingSpot() (parkingSpot, bool) {

	for _, ps := range twm.listOfParkingSpots {
		if ps.isSpotAvailable() {
			return ps, true
		}
	}

	return nil, false
}

func (twm *fourWheelerParkingSpotManager) parkVehicle(v vehicle) (*ticket, bool) {

	ps, ok := twm.findParkingSpot()

	if !ok {
		fmt.Println("Parking Spot is not available.")
		return nil, false
	}
	vehicleSuccessfullyParked := ps.parkVehicle(v)

	if !vehicleSuccessfullyParked {
		return nil, false
	}

	parkedTicket := newTicket(ps, v)
	return parkedTicket, true
}

func (twm *fourWheelerParkingSpotManager) removeVehicle(t *ticket) {
	ps := t.parkingSpot
	ps.removeVehicle()
}

type twoWheelerParkingSpotManager struct {
	listOfParkingSpots []parkingSpot
}

func (twm *twoWheelerParkingSpotManager) getVehicleTypeForParkingSpotManager() string {
	return "Two Wheeler"
}

func (twm *twoWheelerParkingSpotManager) findParkingSpot() (parkingSpot, bool) {

	for _, ps := range twm.listOfParkingSpots {
		if ps.isSpotAvailable() {
			return ps, true
		}
	}

	return nil, false
}

func (twm *twoWheelerParkingSpotManager) addParkingSpot(ps parkingSpot) {
	twm.listOfParkingSpots = append(twm.listOfParkingSpots, ps)
}

func (twm *twoWheelerParkingSpotManager) removeParkingSpot(ps parkingSpot) {
	for index, lps := range twm.listOfParkingSpots {

		if ps.getParkingSpotId() == lps.getParkingSpotId() {
			twm.listOfParkingSpots = append(twm.listOfParkingSpots[0:index], twm.listOfParkingSpots[index+1:]...)
		}
	}
}

func (twm *twoWheelerParkingSpotManager) parkVehicle(v vehicle) (*ticket, bool) {

	ps, ok := twm.findParkingSpot()

	if !ok {
		fmt.Println("Parking Spot is not available.")
		return nil, false
	}
	vehicleSuccessfullyParked := ps.parkVehicle(v)

	if !vehicleSuccessfullyParked {
		return nil, false
	}

	parkedTicket := newTicket(ps, v)
	return parkedTicket, true
}

func (twm *twoWheelerParkingSpotManager) removeVehicle(t *ticket) {
	ps := t.parkingSpot
	ps.removeVehicle()
}
