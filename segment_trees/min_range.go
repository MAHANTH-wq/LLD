package main

const (
	INT_MAX = 1000000007
)

type minRangeSegmentTree struct {
	seg []int // array for storing segment values
	n   int   // no of elements for which the segment tree is being constructed.
	arr []int // array of elements
}

func newMinRangeSegmentTree(arr []int, size int) *minRangeSegmentTree {
	return &minRangeSegmentTree{
		seg: make([]int, 4*size+1),
		n:   size,
		arr: arr,
	}
}

func (ms *minRangeSegmentTree) Build(index int, low int, high int) int {

	if low == high {
		ms.seg[index] = ms.arr[low]
		return ms.seg[index]
	}

	mid := (low + high) / 2

	left := ms.Build(2*index+1, low, mid)
	right := ms.Build(2*index+2, mid+1, high)

	ms.seg[index] = min(left, right)

	return ms.seg[index]
}

func (ms *minRangeSegmentTree) Query(index int, low int, high int, l int, r int) int {

	//No Overlap
	if low > r || high < l {
		return INT_MAX
	}

	//Complete Overlap case
	if low >= l && high <= r {
		return ms.seg[index]
	}

	//Partial Overlap

	mid := (low + high) / 2
	leftValue := ms.Query(2*index+1, low, mid, l, r)
	rightValue := ms.Query(2*index+2, mid+1, high, l, r)

	return min(leftValue, rightValue)

}

func (ms *minRangeSegmentTree) PointUpdate(arrIndex int, value int, segIndex int, low int, high int) {

	//Find the node to update
	if low == high {
		ms.arr[arrIndex] = value
		ms.seg[segIndex] = value
		return
	}

	mid := (low + high) / 2

	if arrIndex <= mid {
		ms.PointUpdate(arrIndex, value, 2*segIndex+1, low, mid)
	} else {
		ms.PointUpdate(arrIndex, value, 2*segIndex+2, mid+1, high)
	}

	ms.seg[segIndex] = min(ms.seg[2*segIndex+1], ms.seg[2*segIndex+2])
	return
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
