package main

import "fmt"

type singletonLazy struct {
	attributeOne   int
	attributeTwo   int
	attributeThree int
}

//In lazy initialization of singleton instance, the instance is created at the time of first object creation
//Lazy initialization is not thread safe.
var singletonLazyInstance *singletonLazy

func newSingletonLazy() *singletonLazy {

	if singletonLazyInstance == nil {
		fmt.Println("Creating new singleton lazy instance")
		singletonLazyInstance = &singletonLazy{
			attributeOne:   1,
			attributeTwo:   2,
			attributeThree: 3,
		}
	}

	return singletonLazyInstance
}

func (s *singletonLazy) printDetails() {
	fmt.Println("Attribute One: ", s.attributeOne)
	fmt.Println("Attribute Two: ", s.attributeTwo)
	fmt.Println("Attribute Three: ", s.attributeThree)

}
