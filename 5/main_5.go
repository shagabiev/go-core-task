package main

import "fmt"

func main() {
	a := []int{64, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(intersectionSlices(a, b))
}

func intersectionSlices(a, b []int) (bool, []int) {
	result := []int{}

	set := make(map[int]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := set[v]; exists {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}
