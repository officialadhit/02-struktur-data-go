package main

import "fmt"

func intersect(a, b []string) []string {
	// map utk mencatat elemen unik dari slice pertama
	seen := make(map[string]bool)
	for _, item := range a {
		seen[item] = true
	}

	result := []string{}
	// map utk mencegah item yang sama lebih dari sekali ke hasil
	added := make(map[string]bool)

	for _, item := range b {
		// jika item ada di slice 'a' dan blm pernah masuk ke result
		if seen[item] && !added[item] {
			result = append(result, item)
			added[item] = true
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]string{"apple", "mango", "banana", "guava"}, []string{"apple", "biscuit", "cereal", "guava"})) // [apple guava]
	fmt.Println(intersect([]string{}, []string{"apple", "banana"}))                                                        // []
	fmt.Println(intersect([]string{"apple", "banana"}, []string{}))                                                        // []
	fmt.Println(intersect([]string{"apple", "banana"}, []string{"mango", "orange"}))                                       // []
	fmt.Println(intersect([]string{"apple", "banana", "mango"}, []string{"mango", "banana"}))                              // [banana mango]
	fmt.Println(intersect([]string{"apple", "apple", "banana"}, []string{"apple", "orange"}))                              // [apple]
}
