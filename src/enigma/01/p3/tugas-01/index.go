package main

import (
	"fmt"
)

func main() {

	var pilihan int

	fmt.Print("Silahkan pilih menu aplikasi:\n\n")
	fmt.Println("1. Menghitung luas dan keliling lingkaran")
	fmt.Println("2. Menghitung luas dan keliling persegi")
	fmt.Println("3. Menghitung luas permukaan dan volume balok")

	fmt.Print("\n\nMasukkan pilihan anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:

		var jari2, luas, keliling float32

		fmt.Printf("Menghitung luas dan keliling lingkaran\n\n")
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&jari2)

		// proses
		luas = 3.14 * jari2 * jari2
		keliling = 2 * 3.14 * jari2

		fmt.Println("Luas dari lingkaran tersebut adalah: ", luas)
		fmt.Println("keliling dari lingkaran tersebut adalah: ", keliling)
	case 2:

		var sisi, luas, keliling float32

		fmt.Printf("Menghitung luas dan keliling persegi\n\n")
		fmt.Print("Masukkan sisi persegi: ")
		fmt.Scan(&sisi)

		// proses
		luas = sisi * sisi
		keliling = 4 * sisi

		fmt.Println("Luas dari persegi tersebut adalah: ", luas)
		fmt.Println("Keliling dari persegi tersebut adalah: ", keliling)
	case 3:

		var panjang, lebar, tinggi, luas, volume float32

		fmt.Printf("Menghitung luas permukaan dan volume balok\n\n")
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)
		fmt.Print("Masukkan tinggi: ")
		fmt.Scan(&tinggi)

		// proses
		luas = 2 * ((panjang * lebar) + (panjang * tinggi) + (lebar * tinggi))
		volume = panjang * lebar * tinggi

		fmt.Println("Luas permukaan balok tersebut adalah: ", luas)
		fmt.Println("Volume balok tersebut adalah: ", volume)

	default:
		fmt.Print("Pilihan anda salah")
	}
}
