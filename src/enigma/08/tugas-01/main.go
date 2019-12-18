package main

import (
	f "fmt"
	"os"
	rg "regexp"
	sc "strconv"
)

type dataMhs struct {
	namaMhs    string
	jurusanMhs string
	umurMhs    int
}

var isInt = rg.MustCompile("^[0-9]+$")

var mhs1 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs2 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs3 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs4 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}
var mhs5 = dataMhs{namaMhs: "", umurMhs: 0, jurusanMhs: ""}

var nama, jurusan string
var umur int

func main() {

	var menu string

	f.Println("--------------------------------------")
	f.Println("Main Menu")
	f.Println("--------------------------------------")
	f.Println("1. Add Mahasiswa")
	f.Println("2. Delete Mahasiswa")
	f.Println("3. View Mahasiswa")
	f.Println("4. Exit")
	f.Println("--------------------------------------")

	for {
		f.Printf("Masukan menu yang dipilih: ")
		f.Scan(&menu)

		if isInt.MatchString(menu) {
			menu, _ := sc.Atoi(menu)
			if menu > 0 && menu < 5 {
				break
			}
		}

		f.Println("Error: menu salah!")
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
		f.Println("Anda berhasil keluar dari aplikasi")
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
		f.Println("--------------------------------------")
		f.Println("Add Mahasiswa")
		f.Println("--------------------------------------")

		for {
			f.Print("Nama (3-20 karakter) : ")
			f.Scan(&nama)

			if len(nama) > 2 && len(nama) < 21 {
				break
			}
			f.Println("Error: nama harus 3-20 karakter")
		}

		for {
			f.Printf("Umur (min 17 tahun)  : ")
			f.Scan(&umur)
			if umur > 16 {
				break
			}
			f.Println("Error: umur minimal 17 tahun")
		}

		for {
			f.Printf("Jurusan (maks 10 karakter) : ")
			f.Scan(&jurusan)

			if len(jurusan) > 1 {
				break
			}
			f.Println("Error: jurusan maksimal 10 karakter")
		}

		simpanMhs(nama, umur, jurusan)
		return
	}

	f.Println("Error: data mahasiswa sudah penuh")
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
		f.Println("Error: data mahasiswa masih kosong")
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

		f.Println()
		f.Println("Success: data mahasiswa yang terakhir masuk berhasil dihapus")
		f.Println(mhs1, mhs2, mhs3, mhs4, mhs5)
		f.Println()
		return
	}

}

func lihatMhs() {

	if totalMhs() > 0 {
		f.Println("1.View by index")
		f.Println("2.View all data")

		var menu string
		for {
			f.Printf("Masukan menu yang dipilih: ")
			f.Scan(&menu)

			if isInt.MatchString(menu) {
				menu, _ := sc.Atoi(menu)
				if menu > 0 && menu < 3 {
					break
				}
			}

			f.Println("Error: menu yang anda masukkan salah")
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
		f.Println("Info: belum ada data mahasiswa")
	}

}

func viewAllMhs() {

	for i := 1; i <= totalMhs(); i++ {
		switch i {
		case 1:
			f.Println(1)
			f.Println("Nama:", mhs1.namaMhs)
			f.Println("Umur:", mhs1.umurMhs)
			f.Println("Jurusan:", mhs1.jurusanMhs)
		case 2:
			f.Println(2)
			f.Println("Nama:", mhs2.namaMhs)
			f.Println("Umur:", mhs2.umurMhs)
			f.Println("Jurusan:", mhs2.jurusanMhs)
		case 3:
			f.Println(3)
			f.Println("Nama:", mhs3.namaMhs)
			f.Println("Umur:", mhs3.umurMhs)
			f.Println("Jurusan:", mhs3.jurusanMhs)
		case 4:
			f.Println(4)
			f.Println("Nama:", mhs4.namaMhs)
			f.Println("Umur:", mhs4.umurMhs)
			f.Println("Jurusan:", mhs4.jurusanMhs)
		case 5:
			f.Println(5)
			f.Println("Nama:", mhs5.namaMhs)
			f.Println("Umur:", mhs5.umurMhs)
			f.Println("Jurusan:", mhs5.jurusanMhs)
		}
	}
}

func viewByIndex() {

	var index int
	for {
		f.Printf("Masukkan index yang ingin ditampilkan: ")
		f.Scan(&index)

		if index == 1 {
			f.Println(1)
			f.Println("Nama:", mhs1.namaMhs)
			f.Println("Umur:", mhs1.umurMhs)
			f.Println("Jurusan:", mhs1.jurusanMhs)
			return
		} else if index == 2 {
			f.Println(2)
			f.Println("Nama:", mhs2.namaMhs)
			f.Println("Umur:", mhs2.umurMhs)
			f.Println("Jurusan:", mhs2.jurusanMhs)
			return
		} else if index == 3 {
			f.Println(3)
			f.Println("Nama:", mhs3.namaMhs)
			f.Println("Umur:", mhs3.umurMhs)
			f.Println("Jurusan:", mhs3.jurusanMhs)
			return
		} else if index == 4 {
			f.Println(4)
			f.Println("Nama:", mhs4.namaMhs)
			f.Println("Umur:", mhs4.umurMhs)
			f.Println("Jurusan:", mhs4.jurusanMhs)
			return
		} else if index == 5 {
			f.Println(5)
			f.Println("Nama:", mhs5.namaMhs)
			f.Println("Umur:", mhs5.umurMhs)
			f.Println("Jurusan:", mhs5.jurusanMhs)
			return
		} else {
			f.Println("Error: anda memasukkan index yang salah")
		}
	}

}
