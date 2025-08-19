package main

import (
	"fmt"
)

// Pipeline represents a 3-stage file processing pipeline
type Pipeline struct {
	files []string
}

// NewPipeline is a constructor
func NewPipeline(files []string) *Pipeline {
	return &Pipeline{files: files}
}

// Stage 1: Source of file names
func (p *Pipeline) sourceFileProcessor() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, file := range p.files {
			out <- file
		}
	}()
	return out
}

// Stage 2: Filter files
func (p *Pipeline) filterFilesProcessor(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for file := range in {
			if len(file) > 5 { // Example condition
				out <- file
			}
		}
	}()
	return out
}

// Stage 3: Process files
func (p *Pipeline) fileProcessor(in <-chan string) {
	for file := range in {
		processedFile := "Processed: " + file
		fmt.Println(processedFile)
	}
}

// Run executes the pipeline
func (p *Pipeline) Run() {
	src := p.sourceFileProcessor()
	filtered := p.filterFilesProcessor(src)
	p.fileProcessor(filtered)
}
