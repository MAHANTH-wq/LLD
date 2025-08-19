package main

import (
	"fmt"
	"sync"
)

type wordProcessor struct {
	listOfWords []string
	prevWord    string
}

func newWordProcessor(words []string) *wordProcessor {
	return &wordProcessor{
		listOfWords: words,
		prevWord:    "",
	}
}

func (wp *wordProcessor) Run() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	sourceChan := wp.sourceWordProcessor(&wg)
	filteredChan := wp.filterWords(sourceChan, &wg)

	wp.printTheWords(filteredChan, &wg)

	wg.Wait()
}

func (wp *wordProcessor) sourceWordProcessor(wg *sync.WaitGroup) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		defer wg.Done()
		for _, word := range wp.listOfWords {
			out <- word
		}
	}()
	return out
}

func (wp *wordProcessor) filterWords(in <-chan string, wg *sync.WaitGroup) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		defer wg.Done()
		for word := range in {
			if word == wp.prevWord {
				continue
			}
			out <- word
			wp.prevWord = word
		}
	}()
	return out
}

func (wp *wordProcessor) printTheWords(in <-chan string, wg *sync.WaitGroup) {

	go func() {
		defer wg.Done()
		for word := range in {
			fmt.Println(word)
		}
	}()

}
