package main

import "fmt"

func main() {
	fmt.Println("new go module")
	careTaker := newCareTaker()
	//Version 1
	originator := newOriginator("a1 version one", "a2 version one")
	mementoOne := originator.createMemento()
	careTaker.addMemento(mementoOne)
	fmt.Println("Printint Version 1 of Originator")
	originator.printState()

	//Version 2
	originator.setAttribute1("a1 version two")
	fmt.Println("Printint Version 2 of Originator")
	originator.printState()
	mementoTwo := originator.createMemento()
	careTaker.addMemento(mementoTwo)

	//Version 3
	originator.setAttribute2("a2 version two")
	fmt.Println("Printint Version 3 of Originator")
	originator.printState()

	previousMemento := careTaker.latestMemento()
	originator.restoreMemento(previousMemento)

	//Should be same as version 2 because that is the latest mementon in care taker
	fmt.Println("Restored originator state which should be same as version 2")
	originator.printState()
}
