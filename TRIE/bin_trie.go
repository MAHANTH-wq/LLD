package main

import (
	"math"
)

type BinaryNode struct {
	links [2]*BinaryNode
}

func NewBinaryNode() *BinaryNode {
	return &BinaryNode{}
}

type BinaryTRIE struct {
	root *BinaryNode
}

func NewBinaryTRIE() *BinaryTRIE {
	return &BinaryTRIE{
		root: NewBinaryNode(),
	}
}

func (bn *BinaryNode) ContainsKey(n int) bool {

	if n < 0 || n > 1 {
		return false
	}

	if bn.links[n] == nil {
		return false
	}

	return true

}

func (bn *BinaryNode) Put(n int) {
	bn.links[n] = NewBinaryNode()
}

func (bn *BinaryNode) Get(n int) *BinaryNode {
	if n < 0 || n > 1 {
		return nil
	}

	return bn.links[n]
}

func (bt *BinaryTRIE) Insert(ele int) {

	binaryNumString := convertIntToBinaryString(ele)
	currentNode := bt.root

	for i := 0; i < 32; i++ {
		value := int(binaryNumString[i] - '0')
		if !currentNode.ContainsKey(value) {
			currentNode.Put(value)
		}
		currentNode = currentNode.Get(value)
	}

}

func (bt *BinaryTRIE) FindMaximumXORWithElement(element int) int {
	binaryNumString := convertIntToBinaryString(element)
	result := 0
	currentNode := bt.root
	for i := 0; i < 32; i++ {
		value := int(binaryNumString[i] - '0')
		oppValue := value ^ 1
		bitPower := int(math.Pow(2, float64(31-i)))

		if currentNode.links[oppValue] != nil {
			currentNode = currentNode.Get(oppValue)
			result = result + bitPower
		} else {
			currentNode = currentNode.Get(value)
		}

	}
	return result

}

func convertIntToBinaryString(n int) string {

	num := ""
	for n > 0 {
		rem := n % 2
		n = n / 2
		if rem == 0 {
			num = "0" + num
		} else {
			num = "1" + num
		}
	}

	for len(num) < 32 {
		num = "0" + num
	}

	return num
}
