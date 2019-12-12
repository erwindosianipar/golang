package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*
		var angka int
		fmt.Print("Masukkan sebuah angka: ")
		fmt.Scan(&angka)
		fmt.Println("Angka yang dimauskkan adalah: ", angka)
	*/

	var kal string
	fmt.Print("Masukan kalimat:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kal = scanner.Text()
	fmt.Println("Kalimat yang anda masukan adalah: ", kal)
}
