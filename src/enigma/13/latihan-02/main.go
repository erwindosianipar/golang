package main

import "fmt"

var pembilang = [12]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

var angka int = 190

func main() {
	terbilang(angka)
}

func terbilang(angka int) {
	if angka < 12 {
		fmt.Println(pembilang[angka])
	} else if angka < 20 {
		fmt.Println(pembilang[angka-10], " belas")
	} else if angka < 100 {
		fmt.Println(pembilang[angka/10], " puluh", pembilang[angka%10])
	} else if angka < 200 {
		fmt.Println(" seratus", pembilang[angka-100])
	} else if angka < 1000 {
		fmt.Println(pembilang[angka/100], " ratus", pembilang[angka%100])
	} else if angka < 2000 {
		fmt.Println(" seribu", pembilang[angka-1000])
	} else if angka < 1000000 {
		fmt.Println(pembilang[angka/1000], " ribu", pembilang[angka%1000])
	} else if angka < 1000000000 {
		fmt.Println(pembilang[angka/1000000], " juta", pembilang[angka%1000000])
	} else if angka < 1000000000000 {
		fmt.Println(pembilang[angka/1000000000], " milyar", pembilang[angka%1000000000])
	} else if angka < 1000000000000000 {
		fmt.Println(pembilang[angka/1000000000000], " trilyun", pembilang[angka%1000000000000])
	}
}
