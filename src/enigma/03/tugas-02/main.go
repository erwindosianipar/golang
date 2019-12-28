package main

import "fmt"

func main() {

	fmt.Printf("Tugas mencetak segitiga dengan perulangan dari inputan Golang\n\n")

	var inputan int
	fmt.Printf("Masukkan jumlah segitiga: ")
	fmt.Scan(&inputan)

	for i := 0; i < inputan; i++ {
		b := "*"
		for j := 0; j < i; j++ {
			b += "*"
		}
		println(b)
	}

}
