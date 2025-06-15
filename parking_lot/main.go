package main

import "fmt"

func main() {

	bike1 := getNewBike("bike1")

	car1 := getNewCar("car1")

	parkLotInstance := getNewParkingLot()

	for i := 0; i < 10; i++ {

		ps := &twoWheelerParkingSpot{
			id:            fmt.Sprintf("%d", i),
			isEmpty:       true,
			parkedVehicle: nil,
		}
		ps4 := &fourWheelerParkingSpot{
			id:            fmt.Sprintf("%d", i),
			isEmpty:       true,
			parkedVehicle: nil,
		}

		parkLotInstance.addParkingSpot(ps)
		parkLotInstance.addParkingSpot(ps4)

	}

	for _, manager := range parkLotInstance.managers {
		fmt.Println(manager.findParkingSpot())
	}

	bike1Ticket, _ := parkLotInstance.parkVehicle(bike1)
	car1Ticket, _ := parkLotInstance.parkVehicle(car1)

	fmt.Println(bike1Ticket.parkingSpot.getParkingSpotId())
	fmt.Println(car1Ticket.parkingSpot.getParkingSpotId())

	for _, manager := range parkLotInstance.managers {
		fmt.Println(manager.findParkingSpot())
	}
	parkLotInstance.removeVehicle(bike1Ticket)
	parkLotInstance.removeVehicle(car1Ticket)

	for _, manager := range parkLotInstance.managers {
		fmt.Println(manager.findParkingSpot())
	}
}
