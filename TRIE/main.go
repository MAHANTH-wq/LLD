package main

import "fmt"

func main() {

	trieData := NewTRIE()
	trieData.Insert("app")
	trieData.Insert("apples")
	trieData.Insert("leo")

	if trieData.Search("apple") {
		fmt.Println("Word apple exist in trie")
	} else {
		fmt.Println("Word apple does not exist in trie")
	}

	if trieData.StartsWith("apple") {
		fmt.Println("There are words that start with apple")
	} else {
		fmt.Println("There are no words that start with apple")
	}

	fmt.Println("Word starts with 'leos' ", trieData.StartsWith("leos"))
}
