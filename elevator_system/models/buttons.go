package models

type button struct {
	value int
}

func createButton(v int) *button {

	return &button{
		value: v,
	}

}

func (b *button) GetButtonValue() int {
	return b.value
}
