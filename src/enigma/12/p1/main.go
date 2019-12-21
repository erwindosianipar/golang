package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("World")

	routine()
	cetak("Hello")
	cetak("World")
}

func cetak(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(1 * time.Second)
	}
}

func routine() {
	fmt.Println("Hello")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("World")
}
