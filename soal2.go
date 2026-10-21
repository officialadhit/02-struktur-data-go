package main

import (
	"fmt"
)

func checkDuplicates(items []string) []string {
	counts := make(map[string]int)
	duplicates := []string{}

	for _, item := range items {
		counts[item]++

		if counts[item] == 2 {
			duplicates = append(duplicates, item)
		}
	}

	return duplicates
}

func main() {
	fmt.Println(checkDuplicates([]string{"apple", "mango", "orange", "orange", "banana", "apple", "guava"}))
	fmt.Println(checkDuplicates([]string{}))
	fmt.Println(checkDuplicates([]string{"kiwi", "mango", "banana"}))
	fmt.Println(checkDuplicates([]string{"kiwi", "kiwi"}))
	fmt.Println(checkDuplicates([]string{"a", "b", "a", "b", "c"}))
	fmt.Println(checkDuplicates([]string{"x", "y", "z", "x", "z", "y", "x"}))
}
