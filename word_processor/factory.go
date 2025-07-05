package main

import "fmt"

var listOfCharacters = make(map[byte]*character, 0)

type characterProcessorFactory struct {
}

func newCharacterProcessorFactory() *characterProcessorFactory {

	return &characterProcessorFactory{}

}

func (wp *characterProcessorFactory) newCharacterObject(l byte) *character {

	if listOfCharacters[l] != nil {
		fmt.Println("Letter alreadycached in factory, returning cached letter")
		return listOfCharacters[l]
	}

	listOfCharacters[l] = newCharacter(l)
	return listOfCharacters[l]

}
