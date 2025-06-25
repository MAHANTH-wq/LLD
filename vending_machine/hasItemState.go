package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hasItem struct {
	vm *vendingMachine
}

func (a *hasItem) requestItem() error {
	fmt.Println("Please Enter the item code to select an item:")

	for _, item := range a.vm.items {
		fmt.Println("Item Name: ", item.itemName, " Item Code: ", item.itemCode)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	itemCodeInput, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid Item Code")
		return err
	}

	checkItemsAdded := false
	for _, item := range a.vm.items {
		if item.itemCode == itemCodeInput {
			a.vm.requestedItems = append(a.vm.requestedItems, item)
			checkItemsAdded = true
		}
	}

	if checkItemsAdded == false {
		fmt.Println("Please enter Valid Item Code")
		return fmt.Errorf("Invalid Item Code")

	}

	return nil
}

func (i *hasItem) removeItem() error {
	fmt.Println("Please select the item code to be removed: ")
	for _, item := range i.vm.requestedItems {
		fmt.Println("Item Name: ", item.itemName)
		fmt.Println("Item Code: ", item.itemCode)
		fmt.Println("Item Cost:", item.itemCost)
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	itemCode, _ := strconv.Atoi(input)

	validItemCode := false
	for index, item := range i.vm.requestedItems {

		if item.itemCode == itemCode {
			validItemCode = true
			i.vm.requestedItems = append(i.vm.requestedItems[0:index], i.vm.requestedItems[index+1:]...)
		}
	}

	if !validItemCode {
		fmt.Println("Invalid Item Code to be removed")
	}

	return nil
}
func (a *hasItem) insertMoney() error {

	totalCost := 0
	for _, item := range a.vm.requestedItems {
		totalCost += item.itemCost
	}

	fmt.Println("Please add total amount: ", totalCost)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	amountGiven, err := strconv.Atoi(input)

	if err != nil {
		fmt.Errorf("Entered Invalid amount")
	}

	fmt.Println("amount receivied ", amountGiven)
	a.vm.currentAmount += amountGiven

	if a.vm.currentAmount >= totalCost {
		fmt.Println("Money is sufficient to process the cart.")
		a.vm.setState(a.vm.hasMoneyState)
	} else {
		pendingAmount := totalCost - a.vm.currentAmount
		fmt.Println("Please remove items or add pending amount ", pendingAmount)
	}

	return nil
}

func (a *hasItem) dispenseItem() error {
	fmt.Println("Please add Money to purchase the items")

	return nil
}
