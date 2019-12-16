package main

import (
	"fmt"
	"strconv"
)

func main() {

	deretAngka := func(angka int) string {
		var output string
		for i := 1; i <= angka*2; i++ {
			if i%2 != 0 {
				output += strconv.Itoa(i) + "\n"
			}
		}
		return output
	}

	fmt.Printf(deretAngka(5))
}
