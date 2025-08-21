package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	arrayOfChopSticks := make([]*semaphore.Weighted, 5)
	for i := 0; i < 5; i++ {
		arrayOfChopSticks[i] = semaphore.NewWeighted(1)
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philisopher(ctx, arrayOfChopSticks, i, 5, &wg)
	}

	wg.Wait()

}

func philisopher(ctx context.Context, arr []*semaphore.Weighted, i int, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	if i%2 == 0 {
		arr[i].Acquire(ctx, 1)
		time.Sleep(3 * time.Second) // Added Delay to verify that deadlock is not occuring
		arr[(i+1)%n].Acquire(ctx, 1)
	} else {
		arr[(i+1)%n].Acquire(ctx, 1)
		time.Sleep(3 * time.Second) // Added Delay to verify that deadlock is not occuring
		arr[i].Acquire(ctx, 1)
	}

	fmt.Println("Philospher ", i, " eating")

	arr[i].Release(1)
	arr[(i+1)%n].Release(1)

}
