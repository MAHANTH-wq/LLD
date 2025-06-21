package models

type displayType int

const (
	floorType displayType = iota
	elevatorType
)

type display interface {
	GetAllButtons() []*button
	GetDisplayType() displayType
	ClickButton(*button) error
}

func createFloorDisplay(floorId int, controller *elevatorController, floors int) display {
	listOfButtons := make([]*button, 0)
	for i := 0; i < floors; i++ {
		listOfButtons = append(listOfButtons, createButton(i))
	}
	return &floorDisplay{
		floorId:            floorId,
		buttons:            listOfButtons,
		displayType:        floorType,
		elevatorController: controller,
	}
}

type floorDisplay struct {
	floorId            int
	elevatorController *elevatorController
	buttons            []*button
	displayType        displayType
}

func (f *floorDisplay) GetAllButtons() []*button {
	return f.buttons
}

func (f *floorDisplay) GetDisplayType() displayType {
	return f.displayType
}

func (f *floorDisplay) ClickButton(b *button) error {
	toFloor := b.value

	request := CreateNewRequest(0, f.floorId, toFloor)
	err := f.elevatorController.receiveNewRequest(request)
	if err != nil {
		return err
	}

	return nil

}
