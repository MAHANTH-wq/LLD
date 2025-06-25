package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idleState struct {
	name string
	vm   *vendingMachine
}

func (i *idleState) insertMoney() error {
	fmt.Println("Please add the items to cart")
	return nil
}

func (i *idleState) removeItem() error {
	fmt.Println("No Items Added to cart")
	return nil
}

func (i *idleState) requestItem() error {

	fmt.Println("Please Enter the item code to select an item:")

	for _, item := range i.vm.items {
		fmt.Println("Item Name: ", item.itemName)
		fmt.Println("Item Code: ", item.itemCode)
		fmt.Println("Item Cost: ", item.itemCost)
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
	for _, item := range i.vm.items {
		if item.itemCode == itemCodeInput {
			i.vm.requestedItems = append(i.vm.requestedItems, item)
			checkItemsAdded = true
		}
	}

	if checkItemsAdded == false {
		fmt.Println("Please enter Valid Item Code")
		return fmt.Errorf("Invalid Item Code")

	} else {
		i.vm.setState(i.vm.hasItem)

	}

	return nil
}

func (i *idleState) dispenseItem() error {
	fmt.Println("Please Add Items before proceeding further")
	return nil
}
