package main

import "fmt"

func main() {

	// ini komentar

	/*
		ini komentar
		multiline
	*/

	var angka int16
	angka = 100

	var kal string
	kal = "ini adalah sebuah string"

	var berkoma float32
	berkoma = 1.5

	var kondisi bool
	kondisi = true
	kondisi = false

	var angka2 = angka

	angka3 := 10
	angka4 := angka3

	kalimat2 := "saya adalah kalimat ke 2"

	berkoma2 := 2.5

	kondisi2 := true

	const phi = 3.14

	angka7, angka8 := 700, "ini adalah string"

	var (
		angka9  int = 9
		angka10 int = angka9
	)

	fmt.Println(angka)
	fmt.Println(angka2)
	fmt.Println(angka3)
	fmt.Println(angka4)
	fmt.Println(angka7)
	fmt.Println(angka8)
	fmt.Println(angka9)
	fmt.Println(angka10)
	fmt.Println(kal)
	fmt.Println(kalimat2)
	fmt.Println(berkoma)
	fmt.Println(berkoma2)
	fmt.Println(kondisi)
	fmt.Println(kondisi2)
}
