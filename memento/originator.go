package main

import "fmt"

type originator struct {
	attribute1 string
	attribute2 string
}

func newOriginator(a1 string, a2 string) *originator {
	return &originator{
		attribute1: a1,
		attribute2: a2,
	}
}
func (o *originator) createMemento() *memento {
	return &memento{
		attribute1: o.getAttribute1(),
		attribute2: o.getAttribute2(),
	}
}

func (o *originator) setAttribute1(v string) {
	o.attribute1 = v
}

func (o *originator) getAttribute1() string {
	return o.attribute1
}

func (o *originator) setAttribute2(v string) {
	o.attribute2 = v
}

func (o *originator) getAttribute2() string {
	return o.attribute2
}

func (o *originator) restoreMemento(m *memento) {
	o.setAttribute1(m.getAttribute1())
	o.setAttribute2(m.getAttribute2())
}

func (o *originator) printState() {
	fmt.Println("Attribute 1: ", o.getAttribute1(), "Attribute 2: ", o.getAttribute2())
}
