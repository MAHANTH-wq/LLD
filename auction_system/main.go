package main

func main() {
	auctionSystem := newAuctionSystem()
	farmerColleagueOne, _ := createNewFarmerColleague(0, "farmer one", auctionSystem)
	farmerColleagueTwo, _ := createNewFarmerColleague(1, "farmer two", auctionSystem)
	businessmanColleagueOne, _ := createNewBusinessManColleague(2, "business one", auctionSystem)

	farmerColleagueOne.update(2000)
	farmerColleagueTwo.update(1000)
	businessmanColleagueOne.update(5000)
}
