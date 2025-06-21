package utils

import (
	"fmt"
	"sync"
)

type HeapType int

const (
	MinHeapType HeapType = iota
	MaxHeapType
)

type Heap interface {
	GetRoot() *Node
	InsertElement(int)
	BuildHeap(*Node)
	GetTopElement() int
	PopElement()
	HeapTraversal(*Node)
	GetSize() int
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func newNode(v int) *Node {

	return &Node{
		value: v,
		left:  nil,
		right: nil,
	}
}

func CreateHeap(heapType HeapType) Heap {

	switch heapType {
	case MinHeapType:
		return &MinHeap{
			heapAccessMutex: sync.Mutex{},
			totalSize:       0,
			Root:            nil,
		}
	case MaxHeapType:
		return &MaxHeap{
			heapAccessMutex: sync.Mutex{},
			totalSize:       0,
			Root:            nil,
		}
	default:
		return nil
	}

}

type MinHeap struct {
	heapAccessMutex sync.Mutex
	totalSize       int
	Root            *Node
}

func (mh *MinHeap) GetRoot() *Node {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	return mh.Root

}

func (mh *MinHeap) GetSize() int {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	return mh.totalSize
}

func (mh *MinHeap) GetTopElement() int {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	return mh.Root.value
}

func (mh *MinHeap) PopElement() {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	mh.Root.value = 1e9
	mh.BuildHeap(mh.Root)
	mh.totalSize = mh.totalSize - 1

}
func (mh *MinHeap) InsertElement(element int) {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	if mh.Root == nil {
		mh.Root = newNode(element)
		mh.totalSize = mh.totalSize + 1
		return
	}

	queue := make([]*Node, 0)
	queue = append(queue, mh.Root)

	for len(queue) > 0 {

		front := queue[0]
		queue = queue[1:]

		if front.value == element {
			return
		}
		if front.left == nil {
			front.left = newNode(element)
			break
		}
		if front.left.value == element {
			return
		}
		queue = append(queue, front.left)

		if front.right == nil {
			front.right = newNode(element)
			break
		}

		if front.right.value == element {
			return
		}

	}

	mh.BuildHeap(mh.Root)
	mh.totalSize = mh.totalSize + 1

}

func (mh *MinHeap) HeapTraversal(root *Node) {

	if root != nil {
		fmt.Println(root.value)
	}

	if root.left != nil {
		mh.HeapTraversal(root.left)
	}

	if root.right != nil {
		mh.HeapTraversal(root.right)
	}

}

func (mh *MinHeap) BuildHeap(root *Node) {

	if root.left != nil {
		mh.BuildHeap(root.left)
	}

	if root.right != nil {
		mh.BuildHeap(root.right)
	}

	if root.right != nil {

		if root.right.value < root.value {
			temp := root.value
			root.value = root.right.value
			root.right.value = temp
		}
	}

	if root.left != nil {

		if root.left.value < root.value {
			temp := root.value
			root.value = root.left.value
			root.left.value = temp
		}
	}
}

type MaxHeap struct {
	heapAccessMutex sync.Mutex
	totalSize       int
	Root            *Node
}

func (mh *MaxHeap) GetRoot() *Node {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	return mh.Root

}

func (mh *MaxHeap) GetSize() int {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	return mh.totalSize
}

func (mh *MaxHeap) GetTopElement() int {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	return mh.Root.value
}

func (mh *MaxHeap) PopElement() {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	mh.Root.value = -1e9
	mh.BuildHeap(mh.Root)
	mh.totalSize = mh.totalSize - 1

}
func (mh *MaxHeap) InsertElement(element int) {
	mh.heapAccessMutex.Lock()
	defer mh.heapAccessMutex.Unlock()
	var mutex sync.Mutex
	mutex.Lock()
	if mh.Root == nil {
		mh.Root = newNode(element)
		mh.totalSize++
		return
	}

	queue := make([]*Node, 0)
	queue = append(queue, mh.Root)

	for len(queue) > 0 {

		front := queue[0]
		queue = queue[1:]

		if front.value == element {
			return
		}

		if front.left == nil {
			front.left = newNode(element)
			break
		}
		if front.left.value == element {
			return
		}
		queue = append(queue, front.left)

		if front.right == nil {
			front.right = newNode(element)
			break
		}

		if front.right.value == element {
			return
		}

		queue = append(queue, front.right)

	}

	mh.BuildHeap(mh.Root)
	mh.totalSize = mh.totalSize + 1
	mutex.Unlock()

}

func (mh *MaxHeap) HeapTraversal(root *Node) {

	if root != nil {
		fmt.Println(root.value)
	}

	if root.left != nil {
		mh.HeapTraversal(root.left)
	}

	if root.right != nil {
		mh.HeapTraversal(root.right)
	}

}

func (mh *MaxHeap) BuildHeap(root *Node) {

	if root.left != nil {
		mh.BuildHeap(root.left)
	}

	if root.right != nil {
		mh.BuildHeap(root.right)
	}

	if root.right != nil {

		if root.right.value > root.value {
			temp := root.value
			root.value = root.right.value
			root.right.value = temp
		}
	}

	if root.left != nil {

		if root.left.value > root.value {
			temp := root.value
			root.value = root.left.value
			root.left.value = temp
		}
	}
}
