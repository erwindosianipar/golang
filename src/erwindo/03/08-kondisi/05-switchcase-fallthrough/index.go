package main

import "fmt"

func main() {

	fmt.Printf("belajar kondisi fallthrough dalam switch case Golang\n\n")

	var point = 6

	switch {
	case point >= 8:
		fmt.Println("Nilai kamu sempurna")
	case (point < 8) && (point > 3):
		fmt.Println("Nilai kamu hampir sempurna")
		fallthrough
		/*
			fungsi fallthough akan mengeksekusi program case di bawahnya
			walaupun case nya bernilai false
		*/
	case point < 5:
		fmt.Println("Kamu harus belajar lagi")
	default:
		fmt.Println("Kamu harus belajar lagi")
	}
}
