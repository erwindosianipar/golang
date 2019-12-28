package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()
	var jenisKelamin string
	var tahunLahir, usia int

	fmt.Printf("Aplikasi pengisian form biodata Golang\n\n")
	fmt.Printf("Masukkan nama anda: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama := scanner.Text()

	if cekPanjangNama(nama) {
		fmt.Printf("Masukkan tahun kelahiran: ")
		fmt.Scan(&tahunLahir)

		if tahunLahir > 1950 {
			usia = year - tahunLahir

			if usia < 17 {
				fmt.Println("Usia anda harus lebih dari 17 tahun")
			} else {
				fmt.Printf("Masukkan jenis kelamin: ")
				fmt.Scan(&jenisKelamin)

				if cekJenisKelamin(jenisKelamin) {
					fmt.Printf("\n\n")
					fmt.Println("Nama anda:", nama)
					fmt.Println("Tahun lahir:", tahunLahir)
					fmt.Println("Jenis Kelamin:", jenisKelamin)
				} else {
					fmt.Println("Hanya boleh diisi dengan PRIA atau WANITA")
				}
			}
		} else {
			fmt.Println("Error: batas tahun kelahiran adalah: 1950")
		}
	} else {
		fmt.Println("Error: nama harus lebih dari 4 dan kurang dari 20 krakter")
	}
}

func cekPanjangNama(nama string) bool {
	if len(nama) > 4 && len(nama) <= 20 {
		return true
	}
	return false
}

func cekJenisKelamin(jenisKelamin string) bool {
	if strings.ToLower(jenisKelamin) == "pria" || strings.ToLower(jenisKelamin) == "wanita" {
		return true
	}
	return false
}
