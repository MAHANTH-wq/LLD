package main

import (
	"time"
)

func main() {

	b := newBathroom()
	for i := 0; i < 15; i++ {
		go func() {

			b.MaleWantsToEnter(i)
		}()

	}

	for i := 0; i < 15; i++ {
		go func() {
			b.FemaleWantsToEnter(i)

		}()
	}

	time.Sleep(1 * time.Minute)

}
