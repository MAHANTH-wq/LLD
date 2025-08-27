package main

import "fmt"

func main() {

	// arr := []int{3, 5, 1, 8, 2, 0}

	// ms := newMinRangeSegmentTree(arr, len(arr))

	// totalMin := ms.Build(0, 0, 5)

	// fmt.Println("Total Minimum in the array", totalMin)

	// rangeMin := ms.Query(0, 0, 5, 2, 4)

	// fmt.Println("Min Element in the range of index 2 and index 4 is ", rangeMin)

	// //point update

	// ms.PointUpdate(5, 7, 0, 0, 5)
	// totalMinAfterUpdate := ms.Query(0, 0, 5, 0, 5)
	// fmt.Println("Total Minimum in the array", totalMinAfterUpdate)

	//Inversion Count Problem

	secondArr := []int{6, 4, 5, 3, 1, 2}
	totalInversionsCount := findInversionsCountInRange(secondArr)
	fmt.Println("Total Inversions count ", totalInversionsCount)

}
