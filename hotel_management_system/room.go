package main

import "fmt"

type roomType int

const (
	singleRoomType roomType = iota
	doubleRoomType
	deluxRoomType
)

type room interface {
	getId() int
	accept(visitor)
}

func getNewRoom(id int, roomType roomType) room {
	switch roomType {
	case singleRoomType:
		return &singleRoom{
			id:       id,
			roomType: "single",
			balcony:  0,
			bedrooms: 1,
		}
	case doubleRoomType:
		return &doubleRoom{
			id:       id,
			roomType: "double",
			balcony:  1,
			bedrooms: 2,
		}
	case deluxRoomType:
		return &deluxSuite{
			id:       id,
			roomType: "delux",
			balcony:  2,
			bedrooms: 3,
		}
	default:
		fmt.Println("Invalid Room Type")
		return nil
	}
}

type singleRoom struct {
	id       int
	roomType string
	balcony  int
	bedrooms int
}

func (s *singleRoom) getId() int {
	return s.id
}

func (s *singleRoom) accept(v visitor) {

	v.visitSingleRoom(s)
}

type doubleRoom struct {
	id         int
	roomType   string
	television string
	balcony    int
	bedrooms   int
	kitchen    int
}

func (d *doubleRoom) getId() int {
	return d.id
}

func (d *doubleRoom) accept(v visitor) {

	v.visitDoubleRoom(d)
}

type deluxSuite struct {
	id         int
	roomType   string
	television string
	balcony    int
	bedrooms   int
	kitchen    int
}

func (ds *deluxSuite) getId() int {
	return ds.id
}

func (ds *deluxSuite) accept(v visitor) {

	v.visitDeluxRoom(ds)
}
