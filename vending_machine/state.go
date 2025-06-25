package main

type state interface {
	insertMoney() error
	requestItem() error
	dispenseItem() error
	removeItem() error
}
