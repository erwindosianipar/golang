package main

import (
	"fmt"
	"strings"
)

func main() {

	name := " Bang Dani bawa makan  "
	// fmt.Println(strings.TrimSpace(name))

	hapusSpasi(name)
	cekAngka1(1)

}

func hapusSpasi(param string) {
	fmt.Println(strings.TrimSpace(param))
}

func cekAngka1(angka int) bool {
	if angka == 1 {
		return true
	} else {
		return false
	}

}
