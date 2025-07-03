package main

import "fmt"

type prototype interface {
	clone() prototype
}

type student struct {
	rollNumber int
	name       string
	age        int
	fatherName string
}

func newStudent(rollNumber int, name string, age int, fatherName string) *student {
	return &student{
		rollNumber: rollNumber,
		name:       name,
		age:        age,
		fatherName: fatherName,
	}
}
func (s *student) clone() *student {
	return newStudent(s.rollNumber, s.name, s.age, s.fatherName)
}

func (s *student) printStudentDetails() {
	fmt.Println("Roll Number: ", s.rollNumber)
	fmt.Println("Name: ", s.name)
	fmt.Println("Age: ", s.age)
	fmt.Println("Father Name: ", s.fatherName)
}
