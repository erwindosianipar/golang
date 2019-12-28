package main

import "fmt"

func isEven(angka int) (genap bool) {
	if angka%2 == 0 {
		genap = true
	} else {
		genap = false
	}
	return
}

func main() {

	isEven := isEven(10)
	fmt.Println(isEven)
}
