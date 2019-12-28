package main

import "fmt"

func main() {

	isEven := func(angka int) bool {
		if angka%2 == 0 {
			return true
		}
		return false
	}

	fmt.Println(isEven(11))
}
