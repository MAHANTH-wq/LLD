package main

import "fmt"

func main() {

	fmt.Println("<--------------- Eager Implementation Results Start ------------------->")
	singletonEager := newSingletonEager()
	singletonEager.printDetails()

	singletonEagerSecond := newSingletonEager()
	singletonEagerSecond.printDetails()

	fmt.Println("<--------------- Eager Implemenation Results End ---------------->")

	fmt.Println("<--------------- Singleton Lazy Implementation Results Start ------------------->")

	go func() {
		singletonLazy := newSingletonLazy()
		singletonLazy.printDetails()
	}()

	go func() {
		singletonLazySecond := newSingletonLazy()
		singletonLazySecond.printDetails()
	}()

	go func() {
		singletonLazyThird := newSingletonLazy()
		singletonLazyThird.printDetails()
	}()

	fmt.Println("<--------------- Singleton Lazy Implemenation Results End ---------------->")
}
