package main

import (
	f "fmt"
	"os"
	rg "regexp"
	sc "strconv"
)

var isInt = rg.MustCompile("^[0-9]+$")

type dataMhs struct {
	namaMhs    string
	umurMhs    int
	jurusanMhs string
}

var mhs []dataMhs

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
	return len(mhs)
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
	if len(mhs) < 5 {
		mhs = append(mhs, dataMhs{
			namaMhs:    nama,
			umurMhs:    umur,
			jurusanMhs: jurusan,
		})

		f.Println()
		f.Println("Success: data mahasiswa berhasil ditambah")
		f.Println(mhs)
		f.Println()
		return
	}
}

func hapusMhs() {

	if totalMhs() < 1 {
		f.Println("Error: data mahasiswa masih kosong")
	} else {
		i := totalMhs()
		mhs = mhs[:i-1]

		f.Println()
		f.Println("Success: data mahasiswa yang terakhir masuk berhasil dihapus")
		f.Println(mhs)
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
	for i := 0; i < totalMhs(); i++ {
		f.Println()
		f.Println(i)
		f.Println("Nama:", mhs[i].namaMhs)
		f.Println("Umur:", mhs[i].umurMhs)
		f.Println("Jurusan:", mhs[i].jurusanMhs)
		f.Println()
	}
}

func viewByIndex() {
	var index string
	for {
		f.Printf("Masukkan index yang ingin ditampilkan: ")
		f.Scan(&index)

		if isInt.MatchString(index) {
			index, _ := sc.Atoi(index)
			if index < totalMhs() {
				f.Println()
				f.Println(index)
				f.Println("Nama:", mhs[index].namaMhs)
				f.Println("Umur:", mhs[index].umurMhs)
				f.Println("Jurusan:", mhs[index].jurusanMhs)
				f.Println()
				break
			} else {
				f.Println("Error: data mahasiswa tidak ada")
			}
		}
		f.Println("Error: anda memasukkan index yang salah")
	}
}
