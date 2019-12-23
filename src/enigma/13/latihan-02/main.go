package main

import "fmt"

var pembilang = [12]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

var angka int = 909

func main() {
	if angka < 12 {
		fmt.Println(pembilang[angka])
	} else if angka < 20 {
		fmt.Println(pembilang[angka-10], "belas")
	} else if angka < 100 {
		fmt.Println(pembilang[angka/10], "puluh", pembilang[angka%10])
	} else if angka < 1000 {
		fmt.Println(pembilang[angka/100], "ratus", pembilang[angka%10])
	}
}
