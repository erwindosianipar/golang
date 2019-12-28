package main

import "fmt"

func main() {
	fmt.Printf("Tugas mencetak segitiga angka dengan perulangan dari inputan Golang\n\n")

	var inputan int
	fmt.Printf("Masukkan jumlah segitiga: ")
	fmt.Scan(&inputan)

	var i int
	for i = inputan; i > 0; i-- {
		var j int
		for j = i; j > 0; j-- {
			print(j)
		}
		println()
	}
}
