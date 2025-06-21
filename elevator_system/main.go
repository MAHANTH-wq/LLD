package main

import (
	models "elevator_system/models"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Working Go code")

	totalFloors := 10

	building := models.GetNewBuildingInstance(totalFloors)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		for j := 9; j >= 7; j-- {
			wg.Add(1)
			go func() {
				defer wg.Done()
				buttons := building.AllFloors[i].GetFloorDisplay().GetAllButtons()
				fmt.Println("Request from floor ", building.AllFloors[i].GetFloorId(), " to floor", buttons[j].GetButtonValue())
				building.AllFloors[i].GetFloorDisplay().ClickButton(buttons[j])
				// building.ReceiveRequestFromFloor(models.CreateNewRequest(i*j, i, j))
			}()
		}
	}

	// "Pause before sending second burst of requests
	time.Sleep(30 * time.Second)

	for i := 0; i < 3; i++ {
		for j := 9; j >= 7; j-- {
			wg.Add(1)
			go func() {
				defer wg.Done()
				buttons := building.AllFloors[j].GetFloorDisplay().GetAllButtons()
				fmt.Println("Request from", building.AllFloors[j].GetFloorId(), " to floor ", buttons[i].GetButtonValue())

				building.AllFloors[j].GetFloorDisplay().ClickButton(buttons[i])
				// building.ReceiveRequestFromFloor(models.CreateNewRequest(i*j, j, i))
			}()
		}
	}

	wg.Wait()
	building.PrintPendingRequests()

}
