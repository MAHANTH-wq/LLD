package main

import (
	"fmt"
)

type hasMoney struct {
	vm *vendingMachine
}

func (hm *hasMoney) requestItem() error {

	fmt.Println("Cannot Request more items before finishing this order")
	return nil
}

func (hm *hasMoney) removeItem() error {

	hm.vm.setState(hm.vm.hasItem)
	hm.vm.removeItem()
	return nil
}

func (hm *hasMoney) insertMoney() error {
	fmt.Println("Money is sufficient to process the cart. Please select dispense or request more items")

	return nil
}

func (hm *hasMoney) dispenseItem() error {

	for _, item := range hm.vm.requestedItems {
		fmt.Println("dispensing....")
		fmt.Println("Please collect the item ", item.itemName)
		hm.vm.currentAmount -= item.itemCost
	}
	fmt.Println("Sending the extra remaining amount: ", hm.vm.currentAmount)
	hm.vm.requestedItems = make([]*item, 0)
	hm.vm.currentAmount = 0

	hm.vm.setState(hm.vm.idleState)
	return nil
}
