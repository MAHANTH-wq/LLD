package main

type ticket struct {
	vehicle     vehicle
	parkingSpot parkingSpot
	price       int
}

func newTicket(ps parkingSpot, v vehicle) *ticket {

	return &ticket{
		vehicle:     v,
		parkingSpot: ps,
		price:       calculateParkingCost(v),
	}
}

func calculateParkingCost(v vehicle) int {

	if v.getVehicleType() == "Two Wheeler" {
		return 10

	} else if v.getVehicleType() == "Four Wheeler" {
		return 100
	} else {
		return 0
	}

}
