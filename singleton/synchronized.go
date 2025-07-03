package main

import (
	"fmt"
	"sync"
)

type singletonSynchronized struct {
	attributeOne   int
	attributeTwo   int
	attributeThree int
}

var singletonSynchronizedInstance *singletonSynchronized
var synchronizedMutex sync.Mutex

func newSingletonSynchronized() *singletonSynchronized {
	synchronizedMutex.Lock()
	defer synchronizedMutex.Unlock()

	if singletonSynchronizedInstance == nil {
		fmt.Println("Creating new singleton synchronized instance")
		singletonSynchronizedInstance = &singletonSynchronized{
			attributeOne:   1,
			attributeTwo:   2,
			attributeThree: 3,
		}
	}

	return singletonSynchronizedInstance
}
