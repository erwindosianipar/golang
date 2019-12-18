package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type dataMhs struct {
	namaMhs    string
	jurusanMhs string
	umurMhs    int
}

var isInt = regexp.MustCompile("^[0-9]+$")

var mhs1 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs2 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs3 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs4 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs5 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}

var nama, jurusan string
var umur int

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
	if mhs1.namaMhs != "" {
		totalMhs++
	}
	if mhs2.namaMhs != "" {
		totalMhs++
	}
	if mhs3.namaMhs != "" {
		totalMhs++
	}
	if mhs4.namaMhs != "" {
		totalMhs++
	}
	if mhs5.namaMhs != "" {
		totalMhs++
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
			fmt.Scan(&nama)

			if len(nama) > 2 && len(nama) < 21 {
				break
			}
			fmt.Println("Error: nama harus 3-20 karakter")
		}

		for {
			fmt.Printf("Umur (min 17 tahun)  : ")
			fmt.Scan(&umur)
			if umur > 16 {
				break
			}
			fmt.Println("Error: umur minimal 17 tahun")
		}

		for {
			fmt.Printf("Jurusan (maks 10 karakter) : ")
			fmt.Scan(&jurusan)

			if len(jurusan) > 1 {
				break
			}
			fmt.Println("Error: jurusan maksimal 10 karakter")
		}

		simpanMhs(nama, umur, jurusan)
		return
	}

	fmt.Println("Error: data mahasiswa sudah penuh")
	return

}

func simpanMhs(nama string, umur int, jurusan string) {
	if mhs1.namaMhs == "" {
		mhs1.namaMhs = nama
		mhs1.umurMhs = umur
		mhs1.jurusanMhs = jurusan
	} else if mhs2.namaMhs == "" {
		mhs2.namaMhs = nama
		mhs2.umurMhs = umur
		mhs2.jurusanMhs = jurusan
	} else if mhs3.namaMhs == "" {
		mhs3.namaMhs = nama
		mhs3.umurMhs = umur
		mhs3.jurusanMhs = jurusan
	} else if mhs4.namaMhs == "" {
		mhs4.namaMhs = nama
		mhs4.umurMhs = umur
		mhs4.jurusanMhs = jurusan
	} else if mhs5.namaMhs == "" {
		mhs5.namaMhs = nama
		mhs5.umurMhs = umur
		mhs5.jurusanMhs = jurusan
	}
}

func hapusMhs() {

	if totalMhs() < 1 {
		fmt.Println("Error: data mahasiswa masih kosong")
	} else {
		i := totalMhs()

		if i == 1 {
			mhs1.namaMhs = ""
			mhs1.umurMhs = 0
			mhs1.jurusanMhs = ""
		} else if i == 2 {
			mhs2.namaMhs = ""
			mhs2.umurMhs = 0
			mhs2.jurusanMhs = ""
		} else if i == 3 {
			mhs3.namaMhs = ""
			mhs3.umurMhs = 0
			mhs3.jurusanMhs = ""
		} else if i == 4 {
			mhs4.namaMhs = ""
			mhs4.umurMhs = 0
			mhs4.jurusanMhs = ""
		} else if i == 5 {
			mhs5.namaMhs = ""
			mhs5.umurMhs = 0
			mhs5.jurusanMhs = ""
		}

		fmt.Println()
		fmt.Println("Success: data mahasiswa yang terakhir masuk berhasil dihapus")
		fmt.Println(mhs1, mhs2, mhs3, mhs4, mhs5)
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

	for i := 1; i <= totalMhs(); i++ {
		switch i {
		case 1:
			fmt.Println(1)
			fmt.Println("Nama:", mhs1.namaMhs)
			fmt.Println("Umur:", mhs1.umurMhs)
			fmt.Println("Jurusan:", mhs1.jurusanMhs)
		case 2:
			fmt.Println(2)
			fmt.Println("Nama:", mhs2.namaMhs)
			fmt.Println("Umur:", mhs2.umurMhs)
			fmt.Println("Jurusan:", mhs2.jurusanMhs)
		case 3:
			fmt.Println(3)
			fmt.Println("Nama:", mhs3.namaMhs)
			fmt.Println("Umur:", mhs3.umurMhs)
			fmt.Println("Jurusan:", mhs3.jurusanMhs)
		case 4:
			fmt.Println(4)
			fmt.Println("Nama:", mhs4.namaMhs)
			fmt.Println("Umur:", mhs4.umurMhs)
			fmt.Println("Jurusan:", mhs4.jurusanMhs)
		case 5:
			fmt.Println(5)
			fmt.Println("Nama:", mhs5.namaMhs)
			fmt.Println("Umur:", mhs5.umurMhs)
			fmt.Println("Jurusan:", mhs5.jurusanMhs)
		}
	}
}

func viewByIndex() {

	var index int
	for {
		fmt.Printf("Masukkan index yang ingin ditampilkan: ")
		fmt.Scan(&index)

		if index == 1 {
			fmt.Println(1)
			fmt.Println("Nama:", mhs1.namaMhs)
			fmt.Println("Umur:", mhs1.umurMhs)
			fmt.Println("Jurusan:", mhs1.jurusanMhs)
			return
		} else if index == 2 {
			fmt.Println(2)
			fmt.Println("Nama:", mhs2.namaMhs)
			fmt.Println("Umur:", mhs2.umurMhs)
			fmt.Println("Jurusan:", mhs2.jurusanMhs)
			return
		} else if index == 3 {
			fmt.Println(3)
			fmt.Println("Nama:", mhs3.namaMhs)
			fmt.Println("Umur:", mhs3.umurMhs)
			fmt.Println("Jurusan:", mhs3.jurusanMhs)
			return
		} else if index == 4 {
			fmt.Println(4)
			fmt.Println("Nama:", mhs4.namaMhs)
			fmt.Println("Umur:", mhs4.umurMhs)
			fmt.Println("Jurusan:", mhs4.jurusanMhs)
			return
		} else if index == 5 {
			fmt.Println(5)
			fmt.Println("Nama:", mhs5.namaMhs)
			fmt.Println("Umur:", mhs5.umurMhs)
			fmt.Println("Jurusan:", mhs5.jurusanMhs)
			return
		} else {
			fmt.Println("Error: anda memasukkan index yang salah")
		}
	}

}
