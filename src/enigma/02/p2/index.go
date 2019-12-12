package main

import "fmt"

func main() {
	fmt.Print("Belajar pengkondisian pada Golang\n\n")

	/* 1. if

	var a int = 15

	fmt.Println("sebelum if")
	if a < 10 {
		fmt.Println("perintah dineksekusi")
	}
	fmt.Println("setelah if")
	*/

	/* 2. if else

	var a int = 15

	fmt.Println("sebelum if")
	if a < 10 {
		fmt.Println("perintah dieksekusi")
	} else {
		fmt.Println("perintah tidak dieksekusi, yang dieksekusi else")
	}
	fmt.Println("setelah if")

	*/

	var a int = 5

	fmt.Println("sebelum if")
	if a < 10 && a > 5 {
		fmt.Println("perintah dieksekusi")
	} else if a == 5 {
		fmt.Println("perintah tidak dieksekusi, yang dieksekusi else if")
	} else {
		fmt.Println("FINAL ELSE")
	}
	fmt.Println("setelah if")
}
