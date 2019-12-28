package main

import "fmt"

func main() {
	fmt.Printf("Belajar seleksi kondisi pada Golang\n\n")

	var point = 10 // ubah point 1 sampai dengan 10

	if point > 9 {
		fmt.Print("Anda lulus dengan nilai sempurna")
	} else if point >= 5 {
		fmt.Printf("Anda lulus dengan nilai %d", point)
	} else if point == 4 {
		fmt.Print("Anda hampir lulus")
	} else {
		fmt.Printf("Anda tidak lulus, nilai anda: %d", point)
	}
}
