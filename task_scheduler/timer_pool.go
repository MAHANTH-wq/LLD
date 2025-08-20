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
	timerPoolInstances sync.Pool
}

func (tp *timerPool) Get(timeout time.Duration) *time.Timer {
	t := tp.timerPoolInstances.Get().(*time.Timer)
	t.Reset(timeout)
	return t
}

func (tp *timerPool) Put(t *time.Timer) {
	t.Stop()
	tp.timerPoolInstances.Put(t)
}

var TimerRecycler TimerPool = newTimerPool()

func newTimerPool() TimerPool {
	return &timerPool{
		timerPoolInstances: sync.Pool{
			New: func() interface{} {
				return time.NewTimer(time.Hour) //Default
			},
		},
	}
}
