package main

import "fmt"

func main() {

	var arrAngka [5][4]int

	for i := 0; i < len(arrAngka); i++ {
		for k := 0; k < len(arrAngka[0]); k++ {
			arrAngka[i][k] = k + i*10
		}
	}

	fmt.Print(arrAngka)
}
