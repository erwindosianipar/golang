package main

import "fmt"

func main() {

	var arrAngka [10]int
	arrAngka[0] = 50
	arrAngka[5] = 150

	fmt.Println(arrAngka)

	var arrKata = [...]string{
		"joe",
		"prana",
		"reza",
	}

	fmt.Println(arrKata)
	fmt.Println("panjang data: ", len(arrKata))

	var ArrAngka2 [10]int
	for i := 0; i < len(ArrAngka2); i++ {
		ArrAngka2[i] = i + 1
	}

	fmt.Println(ArrAngka2)

	// for range pada array
	for i := range ArrAngka2 {
		fmt.Printf("array index ke %v nilainya %v", i, ArrAngka2[i])
		fmt.Println()
	}

	// for range pada array jika range array ingin ditampung
	for i, val := range ArrAngka2 {
		fmt.Printf("array index ke %v nilainya %v", i, val)
		fmt.Println()
	}

	// for _ range pada array digunakan jika program tidak memerlukan counter dicetak
	for _, val := range ArrAngka2 {
		fmt.Printf("nilainya %v", val)
		fmt.Println()
	}

}
