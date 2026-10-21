package main

import (
	"fmt"
	"strings"
)

func calculateFreqs(word string) map[string]int {
	result := make(map[string]int)

	// jika string kosong, lgsg kembali map kosong
	if strings.TrimSpace(word) == "" {
		return result
	}

	// memisahkan string berdasarkan spasi menjadi slice kata
	words := strings.Fields(word)

	// menghitung frekuensi setiap kata
	for _, w := range words {
		result[w]++
	}

	return result
}

func main() {
	fmt.Println(calculateFreqs("i love vibe coding"))
	fmt.Println(calculateFreqs(""))
	fmt.Println(calculateFreqs("go go go"))
	fmt.Println(calculateFreqs("red blue red"))
	fmt.Println(calculateFreqs("one two one two"))
	fmt.Println(calculateFreqs("hello world hello go"))
}
