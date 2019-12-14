package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var namaMhs, jurusanMhs string
var umurMhs int

var isInt = regexp.MustCompile("^[0-9]+$")

func main() {

	var menu string

	fmt.Println("--------------------------------------")
	fmt.Println("Main Menu")
	fmt.Println("--------------------------------------")
	fmt.Println("1. Add Mahasiswa")
	fmt.Println("2. Delete Mahasiswa")
	fmt.Println("3. View Mahasiswa")
	fmt.Println("4. Exit")

	for {
		fmt.Printf("Masukan menu yang dipilih: ")
		fmt.Scan(&menu)

		if isInt.MatchString(menu) {
			menu, _ := strconv.Atoi(menu)
			if menu > 0 && menu < 5 {
				break
			}
		}

		fmt.Println("Error: menu salah!")
	}

	switch {
	case menu == "1":
		addMahasiswa()
	}

}

func addMahasiswa() (string, int, string) {

	fmt.Println("--------------------------------------")
	fmt.Println("Add Mahasiswa")
	fmt.Println("--------------------------------------")

	for {
		fmt.Print("Nama (3-20 karakter) : ")
		fmt.Scan()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		namaMhs = scanner.Text()

		if len(namaMhs) > 2 && len(namaMhs) < 21 {
			break
		}
		fmt.Println("Error: nama harus 3-20 karakter")
	}

	for {
		fmt.Printf("Umur (min 17 tahun)  : ")
		fmt.Scan(&umurMhs)
		if umurMhs > 16 {
			break
		}
		fmt.Println("Error: umur minimal 17 tahun")
	}

	for {
		fmt.Printf("Jurusan (maks 10 karakter) : ")
		fmt.Scan()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		jurusanMhs = scanner.Text()

		if len(jurusanMhs) > 1 {
			break
		}
		fmt.Println("Error: jurusan maksimal 10 karakter")
	}

	return namaMhs, umurMhs, jurusanMhs

}
