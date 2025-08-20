package main

import "math/rand"

type taskType int

const (
	normalTaskType taskType = iota
)

type Task interface {
	GetId() int
	GetName() string
}

func NewTask(id int, name string, tType taskType) Task {
	switch tType {
	case normalTaskType:
		return newNormalTask(id, name)
	default:
		panic("Unknown task type")
	}
}

type normalTask struct {
	id   int
	name string
}

func newNormalTask(id int, name string) *normalTask {
	return &normalTask{
		id:   id,
		name: name,
	}
}

func (nt *normalTask) GetId() int {
	return nt.id
}

func (nt *normalTask) GetName() string {
	return nt.name
}

func createNewRandomTask() Task {
	id := rand.Intn(1000)
	name := newRandomString(10)
	return NewTask(id, name, normalTaskType)
}

func newRandomString(n int) string {
	arr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	returnString := make([]rune, n)
	for i := 0; i < n; i++ {
		returnString[i] = arr[rand.Intn(len(arr))]
	}
	return string(returnString)
}
