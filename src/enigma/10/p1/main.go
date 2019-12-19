package main

import "fmt"

type segitiga struct {
	alas   int
	tinggi int
}

func (r segitiga) luasSegitiga(param int) int {
	return r.alas * r.tinggi / param
}

func main() {
	var alas, tinggi int
	fmt.Printf("Input alas: ")
	fmt.Scan(&alas)
	fmt.Printf("Input tinggi: ")
	fmt.Scan(&tinggi)

	var s1 = segitiga{alas, tinggi}
	var s2 = segitiga{1, 4}
	fmt.Println(s1.luasSegitiga(2))
	fmt.Println(s2.luasSegitiga(4))
}
