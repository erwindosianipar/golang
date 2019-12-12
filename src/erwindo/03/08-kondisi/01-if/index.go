package main

import "fmt"

func main() {
	fmt.Print("Belajar seleksi kondisi if pada Golang\n\n")

	var point = 8 // ubah point

	if point >= 6 {
		/*
			jika variabel point memiliki nilai lebih dari atau sama dengan 6
			maka perintah dibawah ini akan dicetak

		*/

		fmt.Printf("Anda lulus dengan nilai: %d\n", point)
	}

	var kondisi = true // ubah menjadi true atau false

	if kondisi {
		fmt.Println("Kondisi bernilai benar")
	}

	if kondisi == false {
		fmt.Println("Kondisi bernilai salah")
	}
}
