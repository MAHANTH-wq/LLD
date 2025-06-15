package main

type parkingLot struct {
	managers            []parkingSpotManager
	allowedVehicleTypes []string
}

func getNewParkingLot() *parkingLot {

	pl := &parkingLot{}
	pl.allowedVehicleTypes = []string{"Two Wheeler", "Four Wheeler"}
	for _, vType := range pl.allowedVehicleTypes {
		pl.managers = append(pl.managers, getNewParkingSpotManager(vType))
	}

	return pl
}

func (pl *parkingLot) addParkingSpot(ps parkingSpot) {

	for _, manager := range pl.managers {

		if manager.getVehicleTypeForParkingSpotManager() == ps.getVehicleType() {
			manager.addParkingSpot(ps)
		}
	}
}

func (pl *parkingLot) removeParkingSpot(ps parkingSpot) {
	for _, manager := range pl.managers {

		if manager.getVehicleTypeForParkingSpotManager() == ps.getVehicleType() {
			manager.removeParkingSpot(ps)
		}
	}
}

func (pl *parkingLot) parkVehicle(v vehicle) (*ticket, bool) {

	for _, manager := range pl.managers {

		if manager.getVehicleTypeForParkingSpotManager() == v.getVehicleType() {
			return manager.parkVehicle(v)
		}
	}

	return nil, false
}

func (pl *parkingLot) removeVehicle(t *ticket) {

	for _, manager := range pl.managers {
		if manager.getVehicleTypeForParkingSpotManager() == t.parkingSpot.getVehicleType() {
			manager.removeVehicle(t)
		}
	}

}
