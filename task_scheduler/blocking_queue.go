package main

import (
	"sync"
	"time"
)

type BlockingQueue interface {

	//Insert Elements to the queue and block the channel if the queue is full
	Put(ele interface{}) bool

	//Insert Elements to the queue and block the channel until timeout if the channel is full
	Offer(ele interface{}, timeout time.Duration) bool

	//Remove and return the first element from the queue, blocking if the queue is empty
	Take() (interface{}, bool)

	//Remove and return the first element from the queue, blocking until timeout if the queue is empty
	Poll(timeout time.Duration) (interface{}, bool)

	//Queue Length
	Len() int

	//Queue Capacity
	Capacity() int
}

type blockingQueue struct {
	buffer        chan interface{}
	capacity      int
	len           int
	timerRecycler TimerPool
	updateLength  sync.Mutex
}

func NewBlockingQueue(capacity int) BlockingQueue {

	if capacity == 0 {
		panic("capacity illegal")
	}

	return &blockingQueue{
		buffer:        make(chan interface{}, capacity),
		capacity:      capacity,
		len:           0,
		timerRecycler: TimerRecycler,
		updateLength:  sync.Mutex{},
	}
}

func (bq *blockingQueue) IncreaseLength() {

	bq.updateLength.Lock()
	defer bq.updateLength.Unlock()

	bq.len++
}

func (bq *blockingQueue) DecreaseLength() {

	bq.updateLength.Lock()
	defer bq.updateLength.Unlock()

	bq.len--
}

func (bq *blockingQueue) Put(ele interface{}) bool {

	bq.buffer <- ele
	bq.IncreaseLength()
	return true
}

func (bq *blockingQueue) Offer(ele interface{}, timeout time.Duration) bool {
	t := bq.timerRecycler.Get(timeout)
	select {
	case bq.buffer <- ele:
		bq.IncreaseLength()
		return true
	case <-t.C:
		bq.timerRecycler.Put(t)
		return false
	}
}

func (bq *blockingQueue) Take() (interface{}, bool) {
	ele := <-bq.buffer
	bq.DecreaseLength()
	return ele, true
}

func (bq *blockingQueue) Poll(timeout time.Duration) (interface{}, bool) {
	t := bq.timerRecycler.Get(timeout)
	select {
	case ele := <-bq.buffer:
		bq.DecreaseLength()
		return ele, true
	case <-t.C:
		bq.timerRecycler.Put(t)
		return nil, false
	}

}

func (bq *blockingQueue) Len() int {
	return bq.len
}

func (bq *blockingQueue) Capacity() int {
	return bq.capacity
}
