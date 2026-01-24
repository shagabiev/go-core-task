package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(difference(slice1, slice2))
}

func difference(slice1, slice2 []string) []string {
	result := make([]string, 0)

	set := make(map[string]struct{})
	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	for _, v := range slice1 {
		if _, exists := set[v]; !exists {
			result = append(result, v)
		}
	}

	return result
}
