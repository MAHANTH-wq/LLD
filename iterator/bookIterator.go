package main

type bookIterator struct {
	index       int
	size        int
	listOfBooks []*book
}

func newBookIterator(books []*book) *bookIterator {
	return &bookIterator{
		index:       -1,
		size:        len(books),
		listOfBooks: books,
	}
}

func (b *bookIterator) hasNext() bool {
	if b.index+1 < b.size {
		return true
	}
	return false
}

func (b *bookIterator) getNext() *book {

	b.index++
	if b.index >= b.size {
		return nil
	}
	return b.listOfBooks[b.index]
}
