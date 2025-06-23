package utils

import "math"

type HashMap struct {
	totalKeys int
	totalSize int
	arr       []*SinglyLinkedList
}

func InitHashMap() *HashMap {

	size := int(math.Pow(2, 4))
	arr := make([]*SinglyLinkedList, size)

	for i := 0; i < size; i++ {
		arr[i] = &SinglyLinkedList{
			head:   nil,
			length: 0,
		}
	}

	return &HashMap{
		totalKeys: 0,
		totalSize: size,
		arr:       arr,
	}
}

func (hm *HashMap) Insert(key int, value int) error {

	index := key % hm.totalSize

	err := hm.arr[index].InsertElement(key, value)

	if err != nil {
		return err
	}

	hm.totalKeys = hm.totalKeys + 1

	return nil
}

func (hm *HashMap) SearchForKey(key int) (int, error) {

	index := key % hm.totalSize

	value, err := hm.arr[index].SearchForKey(key)

	return value, err

}
