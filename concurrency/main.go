package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// wg := sync.WaitGroup{}
	// wg.Add(3)
	// go sourceFileProcessor(ch1, &wg)
	// go filterFilesProcessor(ch1, ch2, &wg)
	// go fileProcessor(ch2, &wg)

	// wg.Wait()

	//Use wait group to ensure all goroutines complete before exiting
	// files1 := []string{
	// 	"LLD/concurrency/file1.go",
	// 	"LLD/concurrency/file1.go",
	// 	"LLD/concurrency/file3.go",
	// 	"LLD/concurrency/file4.go",
	// 	"LLD/concurrency/file4.go",
	// }

	// wp := newWordProcessor(files1)
	// wp.Run()
	// //
	// Use the Rover example
	// rover := newRover()
	// rover.startRover()
	// rover.moveLeft()
	// time.Sleep(1 * time.Second)
	// rover.moveRight()
	// time.Sleep(1 * time.Second)
	// rover.stop()
	// time.Sleep(5 * time.Second)

	// Err Groups Example
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < 10; i++ {
		g.Go(func() error {
			return task(ctx, i)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error occurred:", err)
		cancel()
	} else {
		fmt.Println("All tasks completed successfully")
	}

}

func sourceFileProcessor(downstream chan string, wg *sync.WaitGroup) {

	listOfFiles := []string{
		"LLD/concurrency/file1.go",
		"LLD/concurrency/file2.go",
		"LLD/concurrency/file3.go",
		"bad",
	}

	for _, file := range listOfFiles {
		downstream <- file
	}
	close(downstream)
	wg.Done()
}

func filterFilesProcessor(upstream chan string, downstream chan string, wg *sync.WaitGroup) {
	for file := range upstream {
		if len(file) > 5 { // Example condition: file name length greater than 10
			downstream <- file
		}
	}
	close(downstream)
	wg.Done()
}

func fileProcessor(upstream chan string, wg *sync.WaitGroup) {
	for file := range upstream {
		// Simulate processing the file
		processedFile := "Processed: " + file
		fmt.Println(processedFile)
	}
	wg.Done()

}

func task(ctx context.Context, i int) error {

	fmt.Println("Task", i, "started")
	select {
	case <-ctx.Done():
		fmt.Println("Task", i, "cancelled")
		return ctx.Err()
	case <-time.After(1 * time.Second):
		if i%4 == 0 {
			fmt.Println("Task", i, "failed")
			return fmt.Errorf("task %d failed", i)
		}
		fmt.Println("Task", i, "completed")

		return nil
	}

}
