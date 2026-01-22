package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := make([]int, 10)

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100) + 1
	}

	fmt.Println("Original slice:", originalSlice)

	even := sliceExample(originalSlice)
	fmt.Println("Even numbers:", even)

	added := addElements(originalSlice, 42)
	fmt.Println("Add 42 in originalSlice:", added)

	copied := copySlice(originalSlice)
	fmt.Println("Copied slice:", copied)

	removed := removeElement(originalSlice, 3)
	fmt.Println("After removing index 3:", removed)
}

func sliceExample(slice []int) []int {
	result := []int{}

	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
}

func addElements(slice []int, element int) []int {
	newSlice := make([]int, len(slice), len(slice)+1)
	copy(newSlice, slice)
	newSlice = append(newSlice, element)

	return newSlice
}

func copySlice(slice []int) []int {
	tmp := make([]int, len(slice))
	copy(tmp, slice)

	return tmp
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return copySlice(slice)
	}

	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}
