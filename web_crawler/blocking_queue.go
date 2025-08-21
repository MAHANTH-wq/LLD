package main

import (
	"sync"
	"time"
)

const (
	timeOutForBQ = 5 * time.Second
)

type BlockingQueue interface {

	//Insert element into the queue and block if the queue is full
	Put(ele interface{}) bool

	//Insert element into the queue and block until timeout if the queue is full
	Offer(ele interface{}, timeout time.Duration) bool

	//Take the first element from the queue, blocking if the queue is empty
	Take() (interface{}, bool)

	//Take the first element from the queue, blocking until timeout if the queue is empty
	Poll(timeout time.Duration) (interface{}, bool)

	//Queue Length
	Len() int

	//Queue Capacity
	Capacity() int
}

type blockingQueue struct {
	buffer       chan interface{}
	capacity     int
	length       int
	timeRecycler TimerPool
	updateLength sync.Mutex
}

func NewBlockingQueue(capacity int) BlockingQueue {
	if capacity <= 0 {
		panic("capacity must be greater than 0")
	}
	return &blockingQueue{
		buffer:       make(chan interface{}, capacity),
		capacity:     capacity,
		length:       0,
		timeRecycler: timeRecycler,
		updateLength: sync.Mutex{},
	}

}

func (b *blockingQueue) IncreaseLength() {
	b.updateLength.Lock()
	defer b.updateLength.Unlock()
	b.length++
}

func (b *blockingQueue) DecreasingLength() {
	b.updateLength.Lock()
	defer b.updateLength.Unlock()
	b.length--
}

func (b *blockingQueue) Put(ele interface{}) bool {
	select {
	case b.buffer <- ele:
		b.IncreaseLength()
		return true
	}
}

func (b *blockingQueue) Offer(ele interface{}, timeout time.Duration) bool {
	t := b.timeRecycler.Get(timeout)
	select {
	case <-t.C:
		b.timeRecycler.Put(t)
		return false
	case b.buffer <- ele:
		b.IncreaseLength()
		b.timeRecycler.Put(t)
		return true
	}
}

func (b *blockingQueue) Take() (interface{}, bool) {
	select {
	case ele := <-b.buffer:
		b.DecreasingLength()
		return ele, true
	}
}

func (b *blockingQueue) Poll(timeout time.Duration) (interface{}, bool) {
	t := b.timeRecycler.Get(timeout)
	select {
	case <-t.C:
		b.timeRecycler.Put(t)
		return nil, false
	case ele := <-b.buffer:
		b.DecreasingLength()
		b.timeRecycler.Put(t)
		return ele, true
	}
}

func (b *blockingQueue) Len() int {
	return b.length
}

func (b *blockingQueue) Capacity() int {
	return b.capacity
}
