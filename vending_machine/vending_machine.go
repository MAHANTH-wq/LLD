package main

import "fmt"

type vendingMachine struct {
	idleState     state
	hasItem       state
	hasMoneyState state

	currentState  state
	currentAmount int

	items          []*item
	requestedItems []*item
}

func initVendingMachine() *vendingMachine {

	vm := &vendingMachine{}

	for i := 1; i <= 10; i++ {
		itemName := fmt.Sprintf("item %d", i)
		vm.items = append(vm.items, &item{itemName: itemName, itemCode: i, itemCost: i * 10})
	}

	idleState := &idleState{
		name: "idle state",
		vm:   vm,
	}

	hasItem := &hasItem{
		vm: vm,
	}

	hasMoney := &hasMoney{
		vm: vm,
	}

	vm.idleState = idleState
	vm.hasItem = hasItem
	vm.hasMoneyState = hasMoney
	vm.currentAmount = 0
	vm.setState(vm.idleState)

	return vm
}

func (vm *vendingMachine) setState(s state) {
	vm.currentState = s
}

func (vm *vendingMachine) insertMoney() {
	vm.currentState.insertMoney()

}

func (vm *vendingMachine) requestItem() {
	vm.currentState.requestItem()

}

func (vm *vendingMachine) dispenseItem() {
	vm.currentState.dispenseItem()

}
