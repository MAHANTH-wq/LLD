package main

import "fmt"

type colleague interface {
	getId() int
	getName() string
	update(int)
	receiveNotification()
}

func createNewFarmerColleague(id int, name string, as *auctionSystem) (*farmerColleague, error) {

	newFarmerColleague := &farmerColleague{
		id:            id,
		name:          name,
		auctionSystem: as,
	}
	err := as.addNewColleague(newFarmerColleague)
	if err != nil {
		return nil, err
	}
	return newFarmerColleague, nil

}

type farmerColleague struct {
	id            int
	name          string
	auctionSystem *auctionSystem
}

func (bc *farmerColleague) update(bidAmount int) {
	bc.auctionSystem.updateBiddingAmount(bc, bidAmount)
}

func (bc *farmerColleague) getId() int {
	return bc.id
}

func (bc *farmerColleague) getName() string {
	return bc.name
}

func (bc *farmerColleague) receiveNotification() {
	fmt.Println(fmt.Sprintf("Colleague Name %s recieved notification on updated bidAmount %d", bc.getName(), bc.auctionSystem.getCurrentBiddingAmount()))
}

func createNewBusinessManColleague(id int, name string, as *auctionSystem) (*businessmanColleague, error) {

	newBusinessColleague := &businessmanColleague{
		id:            id,
		name:          name,
		auctionSystem: as,
	}
	err := as.addNewColleague(newBusinessColleague)
	if err != nil {
		return nil, err
	}
	return newBusinessColleague, nil

}

type businessmanColleague struct {
	id            int
	name          string
	auctionSystem *auctionSystem
}

func (bc *businessmanColleague) update(bidAmount int) {
	bc.auctionSystem.updateBiddingAmount(bc, bidAmount)
}

func (bc *businessmanColleague) getId() int {
	return bc.id
}

func (bc *businessmanColleague) getName() string {
	return bc.name
}

func (bc *businessmanColleague) receiveNotification() {
	fmt.Println(fmt.Sprintf("Colleague Name %s recieved notification on updated bidAmount %d", bc.getName(), bc.auctionSystem.getCurrentBiddingAmount()))
}
