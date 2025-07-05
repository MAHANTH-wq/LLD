package main

import "fmt"

type auctionSystem struct {
	currentBiddingAmount int
	colleagues           []colleague
}

func newAuctionSystem() *auctionSystem {
	return &auctionSystem{
		currentBiddingAmount: 100,
		colleagues:           make([]colleague, 0),
	}
}
func (as *auctionSystem) getCurrentBiddingAmount() int {
	return as.currentBiddingAmount
}

func (as *auctionSystem) addNewColleague(colleague colleague) error {
	for _, c := range as.colleagues {
		if c.getId() == colleague.getId() {
			return fmt.Errorf("Colleague with same id already exists")
		}
	}
	as.colleagues = append(as.colleagues, colleague)
	return nil
}
func (as *auctionSystem) updateBiddingAmount(colleague colleague, bidAmount int) {

	if as.currentBiddingAmount >= bidAmount {
		fmt.Println(fmt.Sprintf("Current Bidding Amount %d ,colleague %s raised %d: INVALID", as.currentBiddingAmount, colleague.getName(), bidAmount))
		return
	}
	as.currentBiddingAmount = bidAmount
	raisedAmountString := fmt.Sprintf("Colleague Name %s raised bid amount to %d", colleague.getName(), as.currentBiddingAmount)
	fmt.Println(raisedAmountString)
	for _, c := range as.colleagues {
		if colleague.getId() == c.getId() {
			continue
		}
		c.receiveNotification()
	}

}
