package main

import "fmt"

func main() {
	fmt.Printf("Belajar switch case pada Golang\n\n")

	var nilai = 4 // nilai 1-5

	switch nilai {
	case 1:
		fmt.Print("Nilai anda: 1")
	case 2, 3, 4:
		fmt.Printf("Nilai anda: %d", nilai)
	case 5:
		fmt.Print("Nilai anda: 5")
	default:
		fmt.Print("Ini adalah nilai default")
	}
}
