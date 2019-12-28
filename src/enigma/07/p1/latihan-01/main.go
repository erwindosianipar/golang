package main

import "fmt"

func isOdd(angka int) (bool, bool) {

	if angka > 0 {
		if angka%2 != 0 {
			return true, true
		}
		return false, true
	}
	return false, false
}

func main() {

	cekGanjil, isValid := isOdd(10)

	if !isValid {
		fmt.Println("Angka tidak dapat dicek genap atau tidak")
	} else {
		if !cekGanjil {
			fmt.Println("Angka genap")
		} else {
			fmt.Println("Angka ganjil")
		}
	}
}
