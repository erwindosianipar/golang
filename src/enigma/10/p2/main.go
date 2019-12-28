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

	var mapS3 = make(map[string]segitiga)
	var s1 = segitiga{4, 2}
	var s2 = segitiga{3, 8}
	var a2 = segitiga{4, 4}

	mapS3["s1"] = s1
	mapS3["s2"] = s2
	mapS3["a1"] = s1
	mapS3["s1"] = a2

	fmt.Println("len :", len(mapS3))
}
