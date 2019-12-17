package main

import (
	f "fmt"
	s "strings"
)

func isOnlyContainXO(karakter string) bool {

	toArr := s.Split(karakter, "")
	for _, char := range toArr {
		if char != "x" && char != "o" {
			return false
		}
	}
	return true
}

func cekXO(karakter string) bool {
	o := s.Count(karakter, "o")
	x := s.Count(karakter, "x")

	if o == x {
		return true
	}
	return false
}

func main() {

	karakter := ""
	for {
		f.Printf("Masukkan karakter: ")
		f.Scan(&karakter)
		if isOnlyContainXO(karakter) {
			break
		}
		f.Println("Error: Karakter yang anda masukkan salah")
	}

	f.Println(cekXO(karakter))
}
