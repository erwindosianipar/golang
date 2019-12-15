package main

import "fmt"

var totalMhs int = 0
var dataMhs [5][3]string

func main() {

	dataMhs[0][0] = "erwindo"
	dataMhs[1][0] = "erwindo"
	dataMhs[2][0] = "erwindo"

	fmt.Println(totalDataMhs())
}

func totalDataMhs() int {
	for i := 0; i < 5; i++ {
		if dataMhs[i][0] == "" {
			//fmt.Println("arr kosong", i)
		} else {
			totalMhs++
		}
	}
	return totalMhs
}
