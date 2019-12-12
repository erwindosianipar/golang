package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Budi"
	names[1] = "Buda"
	names[2] = "Budo"

	fmt.Println(names[0], names[1], names[2])

	var fruits = [3]int{1, 2, 3}
	// var fruits = [...]int{1, 2, 3}

	fmt.Println(fruits)
	fmt.Println("Panjang array", len(fruits))
}
