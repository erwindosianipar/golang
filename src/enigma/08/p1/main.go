package main

import "fmt"

func main() {

	var numberA int
	var numberB *int
	var numberC *int
	var numberD *int

	numberA = 10
	numberB = &numberA
	numberC = &*numberB
	numberD = &*numberC

	*numberB = 29
	*numberB = 30

	fmt.Println(numberA)
	fmt.Println(*numberB)
	fmt.Println(*numberC)
	fmt.Println(*numberD)
}
