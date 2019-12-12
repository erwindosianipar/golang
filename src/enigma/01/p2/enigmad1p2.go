package main

import "fmt"

func main() {

	a := 10
	b := 10 + 5
	c := b - a
	d := b * c
	e := 7 % 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	// e++
	// e = e+1
	e += 10
	c = e + 1 + b - 1*2/5
	fmt.Println(e)
	fmt.Println(c)

	condLeft := true
	condRight := false

	condResult := condLeft && condRight

	fmt.Println("cl", condLeft)
	fmt.Println("cr", condRight)
	fmt.Println("res", condResult)

}
