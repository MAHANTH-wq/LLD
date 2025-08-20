package main

import (
	"fmt"
	"image"
	"time"
)

type command int

const (
	left command = iota
	right
	stop
)

type Rover struct {
	position image.Point
	cmd      chan command
}

func newRover() *Rover {
	return &Rover{
		position: image.Point{X: 0, Y: 0},
	}
}

func (r *Rover) startRover() {
	r.cmd = make(chan command)
	go r.drive()
}

func (r *Rover) moveLeft() {
	r.cmd <- left
}

func (r *Rover) moveRight() {
	r.cmd <- right
}

func (r *Rover) stop() {
	r.cmd <- stop
	close(r.cmd)
}

func (r *Rover) drive() {

	updateAfter := time.After(2 * time.Second)

	for {
		select {
		case cmd := <-r.cmd:
			switch cmd {
			case left:
				fmt.Println("Moving Left from position:", r.position)
				r.position.X--
			case right:
				fmt.Println("Moving Right from position:", r.position)
				r.position.X++
			case stop:
				fmt.Println("Stopping at position:", r.position)
				return
			}
		case <-updateAfter:
			fmt.Println("Current position:", r.position)
			updateAfter = time.After(2 * time.Second)
		}
	}
}
