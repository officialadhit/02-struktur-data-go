package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Barang struct {
	Nama        string
	HargaSatuan int
	Jumlah      int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	fmt.Print("masukkan jumlah barang: ")
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &n)
	}

	daftarBarang := make([]Barang, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Println()
		}

		var b Barang

		fmt.Print("masukkan nama barang: ")
		if scanner.Scan() {
			b.Nama = strings.TrimSpace(scanner.Text())
		}

		fmt.Print("masukkan harga satuan: ")
		if scanner.Scan() {
			fmt.Sscanf(scanner.Text(), "%d", &b.HargaSatuan)
		}

		fmt.Print("masukkan jumlah barang: ")
		if scanner.Scan() {
			fmt.Sscanf(scanner.Text(), "%d", &b.Jumlah)
		}

		daftarBarang[i] = b
	}

	fmt.Println()
	fmt.Println("==STRUK ANDA==")
	total := 0

	for _, item := range daftarBarang {
		subtotal := item.HargaSatuan * item.Jumlah
		total += subtotal
		fmt.Printf("%s %dx %d = %d\n", item.Nama, item.Jumlah, item.HargaSatuan, subtotal)
	}

	fmt.Println("==============")
	fmt.Printf("total: %d\n", total)
}
