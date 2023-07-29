package main

import (
	"fmt"
	"learning/helpers"
)

func main() {
	//input := []int{1, 99, 200, 10, 20, 4, 9, 47, 27, 2}
	// fmt.Println("Are thre any duplicates ?", helpers.ContainsDuplicate(input))
	// fmt.Println("Your searched element is on index", helpers.BinarySearch(input, 47))
	fmt.Println("Anagram", helpers.Anagram())
}
