package main

type memento struct {
	attribute1 string
	attribute2 string
}

func (o *memento) getAttribute1() string {
	return o.attribute1
}

func (o *memento) getAttribute2() string {
	return o.attribute2
}
