package main

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningChargeIntoComputer(mac)

	windows := &windows{}
	windowsAdapter := &adapter{
		windows: windows,
	}

	client.insertLightningChargeIntoComputer(windowsAdapter)
}
