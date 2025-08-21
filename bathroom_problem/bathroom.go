package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type gender int

const (
	male gender = iota
	female
)

type bathroom struct {
	counter               int
	lastGender            gender
	oppGenderWaitingTimes int
	mu                    sync.Mutex
	maleWaiters           *list.List
	femaleWaiters         *list.List
}

func newBathroom() *bathroom {

	return &bathroom{
		counter:               0,
		lastGender:            male,
		oppGenderWaitingTimes: 0,
		mu:                    sync.Mutex{},
		maleWaiters:           list.New(),
		femaleWaiters:         list.New(),
	}

}

type waiter struct {
	id     int
	gender gender
	ready  chan<- struct{}
}

func (b *bathroom) MaleWantsToEnter(id int) error {
	// fmt.Println("Male with id ", id, " wants to enter")
	b.mu.Lock()
	if b.counter >= 0 && b.counter <= 2 {
		b.counter++
		b.lastGender = male
		b.notifyAll()
		fmt.Println("Male with id", id, "entered the bathroom counter value ", b.counter)
		b.mu.Unlock()
	} else {
		ready := make(chan struct{})
		maleWaiter := waiter{id: id, gender: male, ready: ready}
		b.maleWaiters.PushBack(maleWaiter)
		b.mu.Unlock()

		select {
		case <-ready:
			b.mu.Lock()
			fmt.Println("Male with id", id, "entered the bathroom counter value ", b.counter)
			b.mu.Unlock()
		}

	}

	time.Sleep(2 * time.Second)
	b.MaleLeavesTheBathroom(id)
	return nil

}

func (b *bathroom) MaleLeavesTheBathroom(id int) {
	b.mu.Lock()
	if b.counter < 0 {
		panic("more than 3 men entered concurrently and leaving the bathroom")
	}

	b.counter--
	fmt.Println("Male ", id, "leaves the bathroom counter value ", b.counter)
	b.lastGender = male
	b.notifyAll()
	b.mu.Unlock()
}

func (b *bathroom) FemaleWantsToEnter(id int) error {
	// fmt.Println("Female with id ", id, " wants to enter")
	b.mu.Lock()
	if b.counter <= 0 && b.counter >= -2 {
		b.counter--
		b.lastGender = female
		b.notifyAll()
		fmt.Println("Female with id", id, " entered the bathroom: counter value ", b.counter)

		b.mu.Unlock()
	} else {
		ready := make(chan struct{})
		femaleWaiter := waiter{id: id, gender: female, ready: ready}
		b.femaleWaiters.PushBack(femaleWaiter)
		b.mu.Unlock()

		select {
		case <-ready:
			b.mu.Lock()
			fmt.Println("Female with id", id, " entered the bathroom counter value ", b.counter)
			b.mu.Unlock()
		}

	}

	time.Sleep(2 * time.Second)
	b.FemaleLeavesTheBathroom(id)
	return nil
}

func (b *bathroom) FemaleLeavesTheBathroom(id int) {

	if b.counter > 0 {
		panic("For females it should not be negative")
	}
	b.mu.Lock()
	b.counter++
	fmt.Println("Female ", id, "leaves the bathroom counter value ", b.counter)
	b.lastGender = female
	b.notifyAll()
	b.mu.Unlock()
}

// NOTE: THIS NOTIFY ALL is executed whenever lock is acquired only in the above functions
func (b *bathroom) notifyAll() {
	for {

		if b.femaleWaiters.Len() == 0 && b.maleWaiters.Len() == 0 {
			break
		}

		if b.counter == 0 {
			b.oppGenderWaitingTimes = 0
			if b.lastGender == male {

				if b.femaleWaiters.Len() > 0 {
					b.allowFemaleWaiter()
				} else if b.maleWaiters.Len() > 0 {
					b.allowMaleWaiter()
				} else {
					break
				}

			} else {
				if b.maleWaiters.Len() > 0 {
					b.allowMaleWaiter()
				} else if b.femaleWaiters.Len() > 0 {
					b.allowFemaleWaiter()
				} else {
					break
				}
			}

		} else if b.counter > 0 && b.counter < 3 {

			if b.oppGenderWaitingTimes > 3 {
				break
			}
			//Already few males are in the bathroom and
			if b.femaleWaiters.Len() > 0 {
				b.oppGenderWaitingTimes++
			}

			if b.maleWaiters.Len() > 0 {
				b.allowMaleWaiter()
			} else {
				break
			}

		} else if b.counter < 0 && b.counter > -3 {

			if b.oppGenderWaitingTimes > 3 {
				break
			}

			if b.maleWaiters.Len() > 0 {
				b.oppGenderWaitingTimes++
			}

			if b.femaleWaiters.Len() > 0 {
				b.allowFemaleWaiter()
			} else {
				break
			}

		} else {
			// fmt.Println("Check why this case occured")
			break
		}

	}
}

// Note this uses shared variable and should be used only when you have mutex
func (b *bathroom) allowFemaleWaiter() {
	nextEle := b.femaleWaiters.Front()
	femaleWaiter := nextEle.Value.(waiter)
	b.lastGender = female
	b.counter--
	b.femaleWaiters.Remove(nextEle)
	close(femaleWaiter.ready)
}

func (b *bathroom) allowMaleWaiter() {

	nextEle := b.maleWaiters.Front()
	maleWaiter := nextEle.Value.(waiter)
	b.lastGender = male
	b.counter++
	b.maleWaiters.Remove(nextEle)
	close(maleWaiter.ready)
}
