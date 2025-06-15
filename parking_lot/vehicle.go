package main

type vehicle interface {
	getVehicleId() string
	getVehicleType() string
}

func getNewBike(id string) vehicle {

	return &bike{
		id:          id,
		vehicleType: "Two Wheeler",
	}
}

func getNewCar(id string) vehicle {
	return &car{
		id:          id,
		vehicleType: "Four Wheeler",
	}
}

type bike struct {
	id          string
	vehicleType string
}

func (b *bike) getVehicleId() string {

	return b.id
}

func (b *bike) getVehicleType() string {
	return "Two Wheeler"
}

type car struct {
	id          string
	vehicleType string
}

func (c *car) getVehicleId() string {

	return c.id
}

func (c *car) getVehicleType() string {
	return "Four Wheeler"
}
