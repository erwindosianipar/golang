package main

import "fmt"

type person struct {
	nama   string
	usia   int
	gender string
}

func main() {

	var person1 person
	person1.nama = "Erwindo"
	person1.usia = 19
	person1.gender = "Male"

	fmt.Println(person1)

}
