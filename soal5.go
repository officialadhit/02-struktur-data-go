package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var queue []string

	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)

		if op == "push" {
			var val string
			fmt.Scan(&val)

			// Cek apakah nilai sudah ada di dalam queue (hanya simpan data unik)
			exists := false
			for _, item := range queue {
				if item == val {
					exists = true
					break
				}
			}

			// Jika belum ada, tambahkan ke queue
			if !exists {
				queue = append(queue, val)
			}
		} else if op == "pop" {
			// Mengambil elemen pertama jika queue tidak kosong (FIFO)
			if len(queue) > 0 {
				queue = queue[1:]
			}
		}
	}

	// Tampilkan hasil output
	fmt.Printf("length: %d\n", len(queue))
	fmt.Printf("items: %s\n", strings.Join(queue, ","))
}
