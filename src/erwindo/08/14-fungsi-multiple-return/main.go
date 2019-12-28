package main

import "fmt"

func multipleReturn(param int) (int, string) {
	angka := param * 2
	kata := "Sukses mengembalikan multiple return"
	return angka, kata
}

func multiReturn() (nama string, usia int) {
	return
}

func main() {

	//fmt.Println(multipleReturn(2))
	fmt.Println(multiReturn())
}
