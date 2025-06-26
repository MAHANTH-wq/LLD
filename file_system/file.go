package main

import "fmt"

type file struct {
	name string
}

func getNewFile(fileName string) *file {
	return &file{
		name: fileName,
	}
}

func (f *file) ls() {
	fmt.Println("FileName: ", f.name)
}
