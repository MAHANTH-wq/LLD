package main

func main() {

	library := newLibraryCollection()
	library.addNewBook(0, "ZeroToOne")
	library.addNewBook(1, "PersonalMBA")
	library.addNewBook(2, "BaghavadGita")
	library.addNewBook(3, "MahaBharatam")

	it := library.createIterator()

	for it.hasNext() {
		it.getNext().printDetails()
	}
}
