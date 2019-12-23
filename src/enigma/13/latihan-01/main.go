package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var kombinasi = []int{1, 5, 7, 9, 11}
var komb = []int{}

var input int

func main() {

	for {
		fmt.Printf("Masukkan angka: ")
		input, _ = strconv.Atoi(scan())

		if !(input > 32) {
			break
		}
		fmt.Println("Error: maksimal 33")
	}

	fmt.Println(cekKombinasi(input))
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func cekKombinasi(param int) []int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if kombinasi[j]+kombinasi[j+1] == param {
				komb = append(komb, kombinasi[i], kombinasi[j+1])
				return komb
			}
		}
	}
	return []int{-1}
}
