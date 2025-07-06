package main

func main() {

	roomCleaningVisitorOne := getNewVisitor(0, roomCleaningVisitorType)
	roomFoodDeliveryVisitorOne := getNewVisitor(1, roomFoodDeliveryVisitorType)

	singleRoomOne := getNewRoom(0, singleRoomType)
	singleRoomTwo := getNewRoom(1, singleRoomType)

	deluxRoomOne := getNewRoom(2, deluxRoomType)

	singleRoomOne.accept(roomCleaningVisitorOne)
	deluxRoomOne.accept(roomFoodDeliveryVisitorOne)
	singleRoomTwo.accept(roomCleaningVisitorOne)
}
