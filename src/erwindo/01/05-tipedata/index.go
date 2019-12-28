package main

import "fmt"

func main() {
	fmt.Println("Pengenalan tipe data pada Golang")

	/*

		// deklarasi variabel
		var bilanganPositif uint8 = 89
		var bilanganNegatif = -123

		fmt.Println(bilanganPositif, "adalah bilangan positif")
		fmt.Println(bilanganNegatif, "adalah bilangan negatif")

	*/

	var bilanganDesimal float32 = 3.14

	fmt.Printf("PHi: %f\n", bilanganDesimal)

	// menampilkan jumlah 2 angka dibelakang koma
	fmt.Printf("PHi: %.2f", bilanganDesimal)

}
