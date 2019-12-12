package main

import "fmt"

func main() {

	fmt.Printf("Tugas mencetak persegi dengan perulangan dari inputan pada Golang\n\n")

	var inputan int
	fmt.Printf("Masukkan jumlah persegi: ")
	fmt.Scan(&inputan)

	for i := 0; i < inputan; i++ {
		var b string
		for j := 0; j < inputan; j++ {
			b += "*"
		}
		fmt.Println(b)
	}

}
