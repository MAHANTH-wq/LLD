package main

import "fmt"

type visitorType int

const (
	roomCleaningVisitorType visitorType = iota
	roomFoodDeliveryVisitorType
)

type visitor interface {
	visitSingleRoom(*singleRoom)
	visitDoubleRoom(*doubleRoom)
	visitDeluxRoom(*deluxSuite)
}

func getNewVisitor(id int, visitorType visitorType) visitor {
	switch visitorType {
	case roomCleaningVisitorType:
		return &roomCleaningVisitor{
			staffId: id,
		}
	case roomFoodDeliveryVisitorType:
		return &roomFoodDeliveryVisitor{
			staffId: id,
		}
	default:
		fmt.Println("Invalid Room Visitor Type")
		return nil

	}
}

type roomCleaningVisitor struct {
	staffId int
}

func (rc *roomCleaningVisitor) visitSingleRoom(room *singleRoom) {
	fmt.Println("Room Cleaning Visited Single Room ", room.getId())
}

func (rc *roomCleaningVisitor) visitDoubleRoom(room *doubleRoom) {
	fmt.Println("Room Cleaning Visited Double Room ", room.getId())
}

func (rc *roomCleaningVisitor) visitDeluxRoom(room *deluxSuite) {
	fmt.Println("Room Cleaning visited delux room suite ", room.getId())
}

type roomFoodDeliveryVisitor struct {
	staffId int
}

func newRoomFoodDeliveryVisitor(id int) *roomFoodDeliveryVisitor {
	return &roomFoodDeliveryVisitor{
		staffId: id,
	}
}

func (rc *roomFoodDeliveryVisitor) visitSingleRoom(room *singleRoom) {
	fmt.Println("Room Food Delivery Visited Single Room ", room.getId())
}

func (rc *roomFoodDeliveryVisitor) visitDoubleRoom(room *doubleRoom) {
	fmt.Println("Room Food Delivery Visited Double Room ", room.getId())
}

func (rc *roomFoodDeliveryVisitor) visitDeluxRoom(room *deluxSuite) {
	fmt.Println("Room Food Delivery Visited delux room suite ", room.getId())
}
