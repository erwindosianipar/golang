package main

import "fmt"

func main() {
	fmt.Printf("Belajar tipe data array pada Golang\n\n")

	// deklarasi array
	var nama [4]string
	nama[0] = "Budi"
	nama[1] = "Budo"
	nama[2] = "Buni"
	nama[3] = "Buno"

	fmt.Println(nama[0])
	fmt.Println(nama)

	var buah = [4]string{"Apple", "Manggo", "Banana", "Melon"}
	fmt.Println(buah)

	// inisialisasi array secara vertikal
	bahasa = [4]string{
		"indonesia",
		"inggris",
		"inggris",
		"inggris",
	}

	fmt.Println(bahasa)
}
