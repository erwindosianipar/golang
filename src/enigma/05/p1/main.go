package main

import "fmt"

func main() {

	// defer 1
	fmt.Printf("A")
	defer fmt.Printf("B")
	fmt.Printf("C")
	defer fmt.Printf("D")
	fmt.Printf("E")

	// defer 2
	a := 1

	defer print(a)
	print(a)
	a = a + 1
	print(a)
}
