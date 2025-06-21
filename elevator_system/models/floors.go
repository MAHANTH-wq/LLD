package models

type floor struct {
	id      int
	display display
}

func createFloor(id int, display display) *floor {
	return &floor{
		id:      id,
		display: display,
	}
}

func (f *floor) GetFloorDisplay() display {
	return f.display
}

func (f *floor) GetFloorId() int {
	return f.id
}
