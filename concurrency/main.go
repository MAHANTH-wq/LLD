package main

import (
	"fmt"
	"sync"
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
	files1 := []string{
		"LLD/concurrency/file1.go",
		"LLD/concurrency/file1.go",
		"LLD/concurrency/file3.go",
		"LLD/concurrency/file4.go",
		"LLD/concurrency/file4.go",
	}

	wp := newWordProcessor(files1)
	wp.Run()

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
