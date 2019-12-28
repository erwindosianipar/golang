package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {

	angka, err1 := strconv.Atoi("25")

	if err1 != nil {
		fmt.Printf("Angka : %v", angka)
	}

	hasil, err2 := bagi(5, 0)
	if err2 != nil {
		log.Fatal("Pembagian gagal karena:", err1.Error())
	}
	fmt.Println("pembagian berhasil, hasilnya :", hasil)
}

func bagi(a, b float32) (float32, error) {
	if b == 0 {
		var err = errors.New("Pembagi tidak boleh 0")
		return 0, err
	}
	return a / b, nil
}
