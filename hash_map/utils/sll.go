package utils

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
}

func NewNode(key int, value int) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  nil,
	}
}

type SinglyLinkedList struct {
	head   *Node
	length int
}

func (sll *SinglyLinkedList) GetLength() int {
	return sll.length
}

func (sll *SinglyLinkedList) SearchForKey(key int) (int, error) {

	if sll.head == nil {
		return 0, fmt.Errorf("No Elements in the array to search for the key")
	}

	start := sll.head

	for start != nil {
		if start.key == key {
			return start.value, nil
		}

		start = start.next
	}

	return 0, fmt.Errorf("No Value for Key %d found", key)

}

func (sll *SinglyLinkedList) InsertElement(key int, value int) error {

	if sll.head == nil {
		sll.head = NewNode(key, value)
		return nil
	}

	var prev *Node
	start := sll.head

	for start != nil {
		if start.key == key {
			start.value = value
			return nil
		}
		prev = start
		start = start.next
	}

	prev.next = NewNode(key, value)

	return nil
}
