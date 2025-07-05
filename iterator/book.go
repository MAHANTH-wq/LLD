package main

import "fmt"

type book struct {
	bookId   int
	bookName string
}

func newBook(id int, name string) *book {
	return &book{
		bookId:   id,
		bookName: name,
	}
}

func (b *book) printDetails() {
	fmt.Println("Book Id ", b.bookId, " Book Name: ", b.bookName)
}
