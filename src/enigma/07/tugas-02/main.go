package main

import (
	f "fmt"
	rg "regexp"
	sv "strconv"
	s "strings"
)

var isInt = rg.MustCompile("^[0-9]+$")
var angkaSlice []int
var sliceAngka []string
var sliceToInt int

func main() {

	angka := ""

	for {
		f.Printf("Masukkan angka: ")
		f.Scan(&angka)
		if isInt.MatchString(angka) {
			if len(angka) > 0 {
				break
			}
		}
		f.Println("Error: inputan harus berupa angka")
	}

	var sliceAngka = s.Split(angka, "")

	for i := 1; i < len(sliceAngka); i++ {
		sliceToInt, _ = sv.Atoi(sliceAngka[i-1] + sliceAngka[i])
		angkaSlice = append(angkaSlice, sliceToInt)
	}
	f.Println(angkaSlice)
	cariAngkaTerbesar(angkaSlice)
}

func cariAngkaTerbesar(sliceAngka []int) {

	var n, terbesar int

	for _, angka := range sliceAngka {
		if angka > n {
			n = angka
			terbesar = n
		}
	}
	f.Println("Angka yang terbesar adalah:", terbesar)
}
