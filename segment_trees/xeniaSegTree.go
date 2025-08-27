package main

// XENIA AND BIT OPERATIONS PROBLEM CODEFORCES

type xeniaSegmentTree struct {
	seg []int // array for storing segment values
	n   int   // total no of elements
	p   int   // n = power(2,p) total no of elements
	arr []int // array of elements
}

func newXeniaSegmentTree(arr []int, p int, size int) *xeniaSegmentTree {
	return &xeniaSegmentTree{
		seg: make([]int, 4*size+1),
		n:   size,
		p:   p,
		arr: arr,
	}
}

func (ms *xeniaSegmentTree) Build() int {

	result := 0
	if ms.p%2 == 0 {
		result = ms.BuildRecur(0, false, 0, ms.n-1)
	} else {
		result = ms.BuildRecur(0, true, 0, ms.n-1)
	}

	return result

}

func (ms *xeniaSegmentTree) BuildRecur(index int, orr bool, low int, high int) int {

	if low == high {
		ms.seg[index] = ms.arr[low]
		return ms.seg[index]
	}

	mid := (low + high) / 2
	left := ms.BuildRecur(2*index+1, !orr, low, mid)
	right := ms.BuildRecur(2*index+2, !orr, mid+1, high)

	if orr {
		ms.seg[index] = left | right
	} else {
		ms.seg[index] = left ^ right
	}

	return ms.seg[index]
}

func (ms *xeniaSegmentTree) Query(l int, r int) int {

	result := -1
	if ms.p%2 == 0 {
		result = ms.QueryRecur(0, false, 0, ms.n-1, l, r)
	} else {
		result = ms.QueryRecur(0, true, 0, ms.n-1, l, r)
	}

	return result
}

func (ms *xeniaSegmentTree) QueryRecur(index int, orr bool, low int, high int, l int, r int) int {

	//No Overlap
	if low > r || high < l {
		return 0
	}

	//Complete Overlap case
	if low >= l && high <= r {
		return ms.seg[index]
	}

	//Partial Overlap

	mid := (low + high) / 2
	leftValue := ms.QueryRecur(2*index+1, !orr, low, mid, l, r)
	rightValue := ms.QueryRecur(2*index+2, !orr, mid+1, high, l, r)

	if orr {
		return leftValue | rightValue
	} else {
		return leftValue ^ rightValue
	}

}

func (ms *xeniaSegmentTree) PointUpdate(arrIndex int, value int) {

	if ms.p%2 == 0 {
		ms.PointUpdateRecur(arrIndex, value, false, 0, 0, ms.n-1)
	} else {
		ms.PointUpdateRecur(arrIndex, value, true, 0, 0, ms.n-1)
	}
}

func (ms *xeniaSegmentTree) PointUpdateRecur(arrIndex int, value int, orr bool, segIndex int, low int, high int) {

	//Find the node to update
	if low == high {
		ms.arr[arrIndex] = value
		ms.seg[segIndex] = value
		return
	}

	mid := (low + high) / 2

	if arrIndex <= mid {
		ms.PointUpdateRecur(arrIndex, value, !orr, 2*segIndex+1, low, mid)
	} else {
		ms.PointUpdateRecur(arrIndex, value, !orr, 2*segIndex+2, mid+1, high)
	}

	if orr {
		ms.seg[segIndex] = ms.seg[2*segIndex+1] | ms.seg[2*segIndex+2]
	} else {
		ms.seg[segIndex] = ms.seg[2*segIndex+1] ^ ms.seg[2*segIndex+2]
	}

	return
}
