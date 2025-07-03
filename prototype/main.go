package main

import "fmt"

func main() {
	fmt.Println("Print Go language")
	studentOne := newStudent(1, "mahanth", 23, "amar")
	studentOne.printStudentDetails()
	studentTwo := studentOne.clone()
	studentTwo.printStudentDetails()

}
