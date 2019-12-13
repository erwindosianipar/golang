package main

import "fmt"

func main() {

	var angka int

	for {
		print("Masukkan angka: ")
		fmt.Scan(&angka)
		if angka > -1 {
			break
		}
	}

	var hasil int
	var output, sim string

	for i := 2; i <= angka*2; i += 2 {
		sim = ""
		if i < angka*2 {
			sim = "+"
		}
		output += fmt.Sprintf("%v %v ", i, sim)
		hasil += i
	}
	fmt.Println(output, "=", hasil)

	var hasil2 int
	for k := 0; k < angka; k++ {
		if k%2 != 0 {
			sim = ""
			if k < angka-2 {
				sim = " +"
			}
			fmt.Print(k, sim, " ")
			hasil2 += k
		}
	}
	fmt.Println("=", hasil2)

}
