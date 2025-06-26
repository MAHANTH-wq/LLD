package main

import "fmt"

func main() {

	file1 := getNewFile("File 1")
	file2 := getNewFile("File 2")
	file3 := getNewFile("File 3")
	file4 := getNewFile("File 4")
	file5 := getNewFile("File 5")

	directory1 := getNewDirectory("directory 1")

	directory2 := getNewDirectory("directory 2")

	directory3 := getNewDirectory("directory 3")

	directory1.addToDirectory(file1)
	directory2.addToDirectory(file2)
	directory2.addToDirectory(file4)
	directory3.addToDirectory(file3)
	directory3.addToDirectory(file5)

	directory1.addToDirectory(directory2)
	directory2.addToDirectory(directory3)

	fmt.Println("All the files in directory 1:")
	directory1.ls()
	fmt.Println("All the files in directory 2:")
	directory2.ls()
	fmt.Println("ALl the files in directory 3:")
	directory3.ls()

}
