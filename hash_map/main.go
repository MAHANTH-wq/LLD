package main

import (
	"fmt"
	"hash_map/utils"
)

func main() {
	hm := utils.InitHashMap()
	hm.Insert(10, 400)
	hm.Insert(12, 200)
	hm.Insert(30, 3)

	val, _ := hm.SearchForKey(12)

	fmt.Println("Value for Key 12 :", val)

}
