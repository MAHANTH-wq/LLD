package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	bq := NewBlockingQueue(5)
	for i := 0; i < 10; i++ {
		g.Go(func() error {
			if i%4 == 0 {
				time.Sleep(2 * time.Second)
			}
			task := createNewRandomTask()
			if !bq.Offer(task, 2*time.Second) {
				return fmt.Errorf("Task Queue is full") // or handle the error as needed
			}
			return nil
		})
	}

	for i := 0; i < 12; i++ {
		g.Go(func() error {
			taskInterface, ok := bq.Poll(4 * time.Second)
			if !ok {
				return fmt.Errorf("Task Queue is empty") // or handle the error as needed
			}
			task, ok := taskInterface.(Task) // Type assertion to Task
			if !ok {
				return fmt.Errorf("Polled item is not of type Task")
			}
			fmt.Printf("Processing task: %s\n", task.GetName())
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		cancel()
		fmt.Println("Error occurred during processing tasks", err)
	} else {
		fmt.Println("All tasks processed successfully")
	}

}
