package main

import "fmt"

func main() {
	fmt.Printf("Belajar kondisi bersarang pada Golang\n\n")

	var nilai = 10
	var point = "A" // ubah nilai menjadi A, B, atau C

	if nilai == 10 {
		if point == "A" {
			fmt.Println("Anda mendapatkan nilai 10 dan point A")
		} else if point == "B" {
			fmt.Println("Anda mendapatkan nilai 10 dan point B")
		} else {
			fmt.Println("Anda mendapatkan nilai 10 dan point C")
		}
	} else {
		fmt.Println("Nilai tidak berhasil untuk melihat point")
	}
}
