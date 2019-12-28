package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var namaMhs, jurusanMhs string
var umurMhs int
var dataMhs [5][3]string

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
	fmt.Println("--------------------------------------")

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

	switch menu {
	case "1":
		tambahMhs()
		main()
	case "2":
		hapusMhs()
		main()
	case "3":
		lihatMhs()
	case "4":
		fmt.Println("Anda berhasil keluar dari aplikasi")
		os.Exit(1)
	}
}

func totalMhs() int {
	totalMhs := 0
	for i := 0; i < 5; i++ {
		if dataMhs[i][0] != "" {
			totalMhs++
		}
	}
	return totalMhs
}

func tambahMhs() {

	if totalMhs() < 5 {
		fmt.Println("--------------------------------------")
		fmt.Println("Add Mahasiswa")
		fmt.Println("--------------------------------------")

		for {
			fmt.Print("Nama (3-20 karakter) : ")
			fmt.Scan(&namaMhs)

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
			fmt.Scan(&jurusanMhs)

			if len(jurusanMhs) > 1 {
				break
			}
			fmt.Println("Error: jurusan maksimal 10 karakter")
		}

		usia := strconv.Itoa(umurMhs)
		simpanMhs(namaMhs, usia, jurusanMhs)
		return
	}

	fmt.Println("Error: data mahasiswa sudah penuh")
	return

}

func simpanMhs(namaMhs, umurMhs, jurusanMhs string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 1; j++ {
			if dataMhs[i][j] == "" {
				dataMhs[i][0] = namaMhs
				dataMhs[i][1] = umurMhs
				dataMhs[i][2] = jurusanMhs

				fmt.Println()
				fmt.Println("Success: data mahasiswa berhasil ditambah")
				fmt.Println(dataMhs)
				fmt.Println()
				return
			}
		}
	}
}

func hapusMhs() {

	if totalMhs() < 1 {
		fmt.Println("Error: data mahasiswa masih kosong")
	} else {
		i := totalMhs()
		dataMhs[i-1][0] = ""
		dataMhs[i-1][1] = ""
		dataMhs[i-1][2] = ""

		fmt.Println()
		fmt.Println("Success: data mahasiswa yang terakhir masuk berhasil dihapus")
		fmt.Println(dataMhs)
		fmt.Println()
		return
	}

}

func lihatMhs() {

	if totalMhs() > 0 {
		fmt.Println("1.View by index")
		fmt.Println("2.View all data")

		var menu string
		for {
			fmt.Printf("Masukan menu yang dipilih: ")
			fmt.Scan(&menu)

			if isInt.MatchString(menu) {
				menu, _ := strconv.Atoi(menu)
				if menu > 0 && menu < 3 {
					break
				}
			}

			fmt.Println("Error: menu yang anda masukkan salah")
		}

		switch menu {
		case "1":
			viewByIndex()
			main()
		case "2":
			viewAllMhs()
			main()
		}
	} else {
		fmt.Println("Info: belum ada data mahasiswa")
	}

}

func viewAllMhs() {

	for i := 0; i < totalMhs(); i++ {
		fmt.Println()
		fmt.Println(i)
		fmt.Println("Nama:", dataMhs[i][0])
		fmt.Println("Umur:", dataMhs[i][1])
		fmt.Println("Jurusan:", dataMhs[i][2])
		fmt.Println()
	}
}

func viewByIndex() {

	var index string
	for {
		fmt.Printf("Masukkan index yang ingin ditampilkan: ")
		fmt.Scan(&index)

		if isInt.MatchString(index) {
			index, _ := strconv.Atoi(index)
			if index < totalMhs() {
				fmt.Println()
				fmt.Println(index)
				fmt.Println("Nama:", dataMhs[index][0])
				fmt.Println("Umur:", dataMhs[index][1])
				fmt.Println("Jurusan:", dataMhs[index][2])
				fmt.Println()
				break
			} else {
				fmt.Println("Error: data mahasiswa tidak ada")
			}
		}
		fmt.Println("Error: anda memasukkan index yang salah")
	}

}
