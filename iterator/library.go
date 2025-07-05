package main

type collection interface {
	createIterator() iterator
}

type libraryCollection struct {
	listOfBooks []*book
}

func newLibraryCollection() *libraryCollection {
	return &libraryCollection{
		listOfBooks: make([]*book, 0),
	}
}

func (lc *libraryCollection) addNewBook(bookId int, bookName string) {

	newBook := newBook(bookId, bookName)
	lc.listOfBooks = append(lc.listOfBooks, newBook)
}

func (lc *libraryCollection) createIterator() iterator {
	return newBookIterator(lc.listOfBooks)
}
