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
}

func routine() {
	fmt.Println("Hello")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("World")
}
