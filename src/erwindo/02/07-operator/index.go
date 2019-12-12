package main

import "fmt"

func main() {
	fmt.Printf("Belajar operator pada Golang\n\n")

	// opearator aritmatika

	nilai1 := 17
	nilai2 := 20

	var hasil int = nilai1 + nilai2
	fmt.Println(nilai1, "+", nilai2, "=", hasil)

	// operator perbandingan

	var value = 4
	var isEqual = (value == 2)

	fmt.Printf("Nilai %d (%t)", value, isEqual)

}
