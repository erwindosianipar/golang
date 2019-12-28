package main

import "fmt"

func main() {
	fmt.Printf("Tugas mencetak segitiga\n\n")

	var inputan int
	fmt.Printf("Masukkan jumlah segitiga: ")
	fmt.Scan(&inputan)

	for k := 0; k < inputan; k++ {
		if k == 0 {
			for i := 0; i < inputan; i++ {
				var b string
				for j := 0; j < i; j++ {
					b += "*"
				}
				println(b)
			}

			for i := inputan; i > 0; i-- {
				var b string
				for j := 0; j < i; j++ {
					b += "*"
				}
				println(b)
			}
		}
		return
	}
}
