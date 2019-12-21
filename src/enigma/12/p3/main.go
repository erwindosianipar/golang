package main

import "fmt"

func jumlah(a, b int) {
	// 1
	fmt.Println(a, b, a+b)
}

func main() {
	// 1
	go jumlah(5, 5) // -> perintah ini tidak akan sempat dicetak karena fungsi main sudah habis di eksekusi
}
