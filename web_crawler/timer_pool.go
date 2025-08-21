package main

import (
	"sync"
	"time"
)

type TimerPool interface {
	Get(time.Duration) *time.Timer
	Put(*time.Timer)
}

type timerPool struct {
	timerPoolInstance sync.Pool
}

func (tp *timerPool) Get(d time.Duration) *time.Timer {
	t := tp.timerPoolInstance.Get().(*time.Timer)
	t.Reset(d)
	return t
}

func (tp *timerPool) Put(t *time.Timer) {
	t.Stop()
	tp.timerPoolInstance.Put(t)
}

var timeRecycler TimerPool = newTimerPool()

func newTimerPool() TimerPool {
	return &timerPool{
		timerPoolInstance: sync.Pool{
			New: func() interface{} {
				return time.NewTimer(time.Hour)
			},
		},
	}
}
