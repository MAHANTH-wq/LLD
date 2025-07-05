package main

func main() {

	letterFactory := newCharacterProcessorFactory()

	letterOne := letterFactory.newCharacterObject('t')

	letterOne.display(10, 33)

	letterTwo := letterFactory.newCharacterObject('t')

	letterTwo.display(10, 33)

}
