package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	fruits := [4]string{"apple", "grape", "banana", "melon"}

	for i, b := range fruits {
		fmt.Println(i, b)
	}

	// jika variabel penampung tidak digunakan, maka digunakan _ agar tidak perlu dicetak
	for _, b := range fruits {
		fmt.Println(b)
	}

	for j := 0; j < len(fruits); j++ {
		fmt.Println(j, fruits[j])
	}
}
