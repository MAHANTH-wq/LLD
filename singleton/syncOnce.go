package main

import "sync"

type singletonUsingSyncOnce struct {
	attributeOne   int
	attributeTwo   int
	attributeThree int
}

var singletonUsingSyncOnceInstance *singletonUsingSyncOnce
var singletonOnce = &sync.Once{}

func newSingletonUsingSyncOnce() *singletonUsingSyncOnce {

	singletonOnce.Do(func() {
		singletonUsingSyncOnceInstance = &singletonUsingSyncOnce{
			attributeOne:   1,
			attributeTwo:   2,
			attributeThree: 3,
		}
	})

	return singletonUsingSyncOnceInstance

}
