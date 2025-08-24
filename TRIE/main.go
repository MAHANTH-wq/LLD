package main

import "fmt"

func main() {

	trieData := NewTRIE()
	trieData.Insert("a")
	trieData.Insert("ap")
	trieData.Insert("app")
	trieData.Insert("app")
	trieData.Insert("appl")
	trieData.Insert("apple")
	trieData.Insert("apples")
	trieData.Insert("apples")
	trieData.Insert("apples")
	trieData.Insert("o")
	trieData.Insert("or")
	trieData.Insert("ora")
	trieData.Insert("oran")
	trieData.Insert("orang")
	trieData.Insert("orange")
	trieData.Insert("orange")
	trieData.Insert("leo")

	fmt.Println("Number of words 'apples' inserted", trieData.Search("apples"))
	fmt.Println("Number of word with 'app' as prefix", trieData.StartsWith("app"))
	fmt.Println("Word starts with 'leos' ", trieData.StartsWith("leos"))

	result := trieData.LongestCompleteString(trieData.root)
	fmt.Println("Longest Complete String", result)

	count := CountNoOfDistinctSubStrings("apple")
	fmt.Println("No of Distinct Substrings of apple ", count)
	count = CountNoOfDistinctSubStrings("ababd")
	fmt.Println("No of Distinct Substrings of ababd ", count)
}
