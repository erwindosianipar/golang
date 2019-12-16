package main

import "fmt"

func multipleReturn(angka int) (int, string) {
	res := angka * 2
	return res, "Sukses mengembalikan multiple return"
}

func main() {

	fmt.Println(multipleReturn(2))
}
