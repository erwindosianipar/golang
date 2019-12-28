package main

import "fmt"

func main() {

	// deklarasi variabel
	var nama, asal string
	var usia int

	fmt.Printf("Masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Printf("Masukkan kota: ")
	fmt.Scan(&asal)
	fmt.Printf("Masukkan usia: ")
	fmt.Scan(&usia)

	fmt.Printf("\n\n")

	fmt.Printf("Nama saya %s, berasal dari kota %s, dan usia saya %d tahun", nama, asal, usia)
}
