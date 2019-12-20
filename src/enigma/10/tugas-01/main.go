package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	rg "regexp"
	sc "strconv"
)

var isInt = rg.MustCompile("^[0-9]+$")

type dataMhs struct {
	NamaMhs    string `json:"nama"`
	UmurMhs    int    `json:"umur"`
	JurusanMhs string `json:"jurusan"`
}

var mhs []dataMhs

var nama, jurusan string
var umur int

var path = "mahasiswa.json"

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
			menu, _ := sc.Atoi(menu)
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
		main()
	case "4":
		fmt.Println("Anda berhasil keluar dari aplikasi")
		os.Exit(0)
	}
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

func totalMhs() int {
	data := readFile()
	if data == "" {
		return 0
	}
	total := jsonDecode(readFile())
	return len(total)
}

func simpanMhs(nama string, umur int, jurusan string) {
	if totalMhs() < 5 {
		var file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		data := readFile()
		mhs = jsonDecode(readFile())
		if data != "[]" && data != "" {
			mhs = append(mhs, dataMhs{NamaMhs: nama, UmurMhs: umur, JurusanMhs: jurusan})
			file.WriteString(jsonEncode(mhs))
		} else {
			var mhs = []dataMhs{{nama, umur, jurusan}}
			file.WriteString(jsonEncode(mhs))
		}
		fmt.Println()
		fmt.Println("Success: data mahasiswa berhasil ditambah")
		fmt.Println()
	}
}

func hapusMhs() {
	if totalMhs() < 1 {
		fmt.Println("Error: data mahasiswa masih kosong")
	} else {
		dataStruct := jsonDecode(readFile())
		data := jsonEncode(dataStruct[:totalMhs()-1])

		os.Remove(path)
		var file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		file.WriteString(data)
		fmt.Println()
		fmt.Println("Success: data mahasiswa yang terakhir masuk berhasil dihapus")
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
				menu, _ := sc.Atoi(menu)
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
	dataStruct := jsonDecode(readFile())
	for i := 0; i < totalMhs(); i++ {
		fmt.Println()
		fmt.Println(i)
		fmt.Println("Nama:", dataStruct[i].NamaMhs)
		fmt.Println("Umur:", dataStruct[i].UmurMhs)
		fmt.Println("Jurusan:", dataStruct[i].JurusanMhs)
		fmt.Println()
	}
}

func viewByIndex() {
	dataStruct := jsonDecode(readFile())
	var index string
	for {
		fmt.Printf("Masukkan index yang ingin ditampilkan: ")
		fmt.Scan(&index)

		if isInt.MatchString(index) {
			index, _ := sc.Atoi(index)
			if index < totalMhs() {
				fmt.Println()
				fmt.Println(index)
				fmt.Println("Nama:", dataStruct[index].NamaMhs)
				fmt.Println("Umur:", dataStruct[index].UmurMhs)
				fmt.Println("Jurusan:", dataStruct[index].JurusanMhs)
				fmt.Println()
				break
			} else {
				fmt.Println("Error: data mahasiswa tidak ada")
			}
		}
		fmt.Println("Error: anda memasukkan index yang salah")
	}
}

func readFile() string {
	var text string
	var file, _ = os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text()
	}
	if text != "[]" && text != "" {
		return text
	}
	return ""
}

func jsonEncode(data []dataMhs) string {
	var jsonData, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	var jsonString = string(jsonData)
	return jsonString
}

func jsonDecode(data string) []dataMhs {
	var user []dataMhs
	var jsonData = []byte(data)
	var err = json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}
