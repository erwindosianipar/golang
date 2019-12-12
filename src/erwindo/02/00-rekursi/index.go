package main

import "fmt"

func main() {
	rekursi(10)
}

func rekursi(nilai int) {
	if nilai == 0 {
		return
	}
	fmt.Println(nilai)
	nilai--

	rekursi(nilai)
}
