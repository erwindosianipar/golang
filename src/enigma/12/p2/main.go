package main

import (
	"fmt"
	"time"
)

func cetak(str string, delay int) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(time.Duration(delay))
	}
}

func main() {
	cetak("1", 100)
	go cetak("2", 300)
	cetak("3", 100)
	cetak("4", 100)
}
