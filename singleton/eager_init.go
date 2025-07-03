package main

import "fmt"

type singletonEager struct {
	attributeOne   int
	attributeTwo   int
	attributeThree int
}

// In eager initialization, the singleton instance is created at the time of declaration, as opposed to lazy initialization, where it is created when first needed.
var singletonEagerInstance = &singletonEager{attributeOne: 1, attributeTwo: 2, attributeThree: 3}

func newSingletonEager() *singletonEager {
	return singletonEagerInstance
}

func (s *singletonEager) printDetails() {
	fmt.Println("Attribute One: ", s.attributeOne)
	fmt.Println("Attribute Two: ", s.attributeTwo)
	fmt.Println("Attribute Three: ", s.attributeThree)
}
