package main

import "fmt"

func main() {
	fmt.Printf("Belajar switch case dengan kondisi pada Golang\n\n")

	var point int = 9 // ubah nilai point 1 sampai dengan 10

	switch {
	case point == 1:
		fmt.Printf("Poin anda adalah: 1")
	case point <= 9:
		fmt.Printf("Poin anda adalah: %d", point)
	case point > 9:
		fmt.Printf("Poin anda sempurna")
	}
}
