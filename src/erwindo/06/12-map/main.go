package main

import "fmt"

func main() {
	fmt.Printf("Belajar map pada Golang\n\n")

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["Januari"] = 50
	chicken["Februari"] = 40

	fmt.Println("Januari", chicken["Januari"])
	fmt.Println("Mei", chicken["Mei"])
}
