package main

type parkingSpot interface {
	getParkingSpotId() string
	isSpotAvailable() bool
	parkVehicle(vehicle) bool
	removeVehicle()
	getVehicleType() string
}

type twoWheelerParkingSpot struct {
	id            string
	isEmpty       bool
	parkedVehicle vehicle
}

func (tw *twoWheelerParkingSpot) getParkingSpotId() string {
	return tw.id
}

func (tw *twoWheelerParkingSpot) isSpotAvailable() bool {
	return tw.isEmpty
}

func (tw *twoWheelerParkingSpot) parkVehicle(v vehicle) bool {
	tw.parkedVehicle = v
	tw.isEmpty = false
	return true
}

func (tw *twoWheelerParkingSpot) removeVehicle() {
	tw.parkedVehicle = nil
	tw.isEmpty = true
}

func (tw *twoWheelerParkingSpot) getVehicleType() string {
	return "Two Wheeler"
}

type fourWheelerParkingSpot struct {
	id            string
	isEmpty       bool
	parkedVehicle vehicle
}

func (fw *fourWheelerParkingSpot) getParkingSpotId() string {
	return fw.id
}

func (fw *fourWheelerParkingSpot) isSpotAvailable() bool {
	return fw.isEmpty
}

func (fw *fourWheelerParkingSpot) parkVehicle(v vehicle) bool {
	fw.parkedVehicle = v
	fw.isEmpty = false
	return true
}

func (fw *fourWheelerParkingSpot) removeVehicle() {
	fw.parkedVehicle = nil
	fw.isEmpty = true
}

func (fw *fourWheelerParkingSpot) getVehicleType() string {
	return "Four Wheeler"
}
