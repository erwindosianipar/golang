package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Aplikasi cek validasi email Golang\n\n")
	fmt.Printf("Masukkan email anda: ")

	var email string
	fmt.Scan(&email)

	// 3.1. validasi tambahan -> cek panjang email 5 - 255 karakter
	if cekPanjangEmail(email) {

		// 1. fungsi validasi ada @
		if adaAt(email) {
			var nameDomain [2]string

			// fungsi split akan membagi / mengiris string
			nameDomain[0] = strings.Split(email, "@")[0]
			nameDomain[1] = strings.Split(email, "@")[1]

			// 3.2. validasi tambahan -> cek email menggunakan domain TLD atau tidak
			if cekDomainTld(nameDomain[1]) {

				// 3.3. validasi tambahan -> cek nama email adalah karakter dalam alphabet
				if cekAlphabet(nameDomain[0]) {

					// 2. fungsi validasi ada '.' pada domain / setelah @
					if adaTitik(nameDomain[1]) {
						fmt.Println("Success: email anda valid")
					} else {
						fmt.Println("Error: domain pada email anda salah")
					}
				} else {
					fmt.Println("Error: email menandung karakter yang tidak valid")
				}
			} else {
				fmt.Println("Error: domain pada email anda tidak TLD")
			}
		} else {
			fmt.Println("Error: email tidak mengandung @")
		}
	} else {
		fmt.Println("Error: email harus lebih dari 5 dan kurang dari 256 karakter")
	}
}

func adaAt(param string) bool {
	// fungsi contains akan mengecek apakah dalam string terdapat karakter yang dicari
	return strings.Contains(param, "@")
}

func adaTitik(param string) bool {
	return strings.Contains(param, ".")
}

func cekAlphabet(param string) bool {
	for _, r := range param {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func cekPanjangEmail(param string) bool {
	if len(param) > 5 && len(param) <= 255 {
		return true
	}
	return false
}

func cekDomainTld(param string) bool {
	domains := [3]string{"com", "org", "net"} // list domain TLD yang diperlukan

	for _, domain := range domains {
		// fungsi dibawah ini akan mengambil panjang string pada belakang karakter
		tld := param[len(param)-3:]
		if tld == domain {
			return true
		}
	}
	return false
}
