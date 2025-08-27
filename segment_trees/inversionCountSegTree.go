package main

type rangeSumSegmentTree struct {
	segArr []int
	n      int
	arr    []int
}

// Assuming all the elements in arr are greater than zero and less than or equal to n
func findInversionsCountInRange(arr []int) int {

	n := len(arr)
	freq := make([]int, n+1)
	for i := 0; i < n; i++ {
		freq[arr[i]]++
	}

	newRSTree := newRangeSumSegmentTree(freq, len(freq))

	newRSTree.Build(0, 0, len(freq)-1)

	count := 0
	// fmt.Println(len(freq))
	for i := 0; i < n; i++ {
		newRSTree.PointUpdate(arr[i], -1, 0, 0, len(freq)-1)
		count = count + newRSTree.Query(0, 0, len(freq)-1, 0, arr[i])
	}

	return count

}

func newRangeSumSegmentTree(arr []int, n int) *rangeSumSegmentTree {

	return &rangeSumSegmentTree{
		segArr: make([]int, 4*n+1),
		n:      n,
		arr:    arr,
	}
}

func (rt *rangeSumSegmentTree) Build(index int, low int, high int) {

	if low == high {
		rt.segArr[index] = rt.arr[low]
		return
	}

	mid := (low + high) / 2
	rt.Build(2*index+1, low, mid)
	rt.Build(2*index+2, mid+1, high)

	rt.segArr[index] = rt.segArr[2*index+1] + rt.segArr[2*index+2]
}

func (rt *rangeSumSegmentTree) Query(index int, low int, high int, l int, r int) int {

	//No overlap
	if low > r || high < l {
		return 0
	}

	// Complete Overlap

	if low >= l && high <= r {
		return rt.segArr[index]
	}

	mid := (low + high) / 2
	leftValue := rt.Query(2*index+1, low, mid, l, r)
	rightValue := rt.Query(2*index+2, mid+1, high, l, r)

	return leftValue + rightValue

}

func (rt *rangeSumSegmentTree) PointUpdate(arrIndex int, value int, segIndex int, low int, high int) {

	if low == high {
		rt.arr[arrIndex] += value
		rt.segArr[segIndex] += value
		return
	}

	mid := (low + high) / 2

	if arrIndex <= mid {
		rt.PointUpdate(arrIndex, value, 2*segIndex+1, low, mid)
	} else {
		rt.PointUpdate(arrIndex, value, 2*segIndex+2, mid+1, high)
	}

	rt.segArr[segIndex] = rt.segArr[2*segIndex+1] + rt.segArr[2*segIndex+2]
	return
}
