package main

import "fmt"

type character struct {
	letter byte
}

func newCharacter(l byte) *character {
	return &character{
		letter: l,
	}
}

//Extrinsic data passed as method paramter

func (c *character) display(row int, column int) {
	//prints the character at row and column level position
	value := fmt.Sprintf("character %c at row %d and column %d", c.letter, row, column)
	fmt.Println(value)
}
