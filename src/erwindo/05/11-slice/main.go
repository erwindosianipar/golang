package main

import "fmt"

func main() {

	var fruitsA = []string{"Apple", "Manggo"}
	var fruitsB = [2]string{"Kiwi", "Banana"}
	var fruitsC = [...]string{"Kiwi", "Banana"}

	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)
}
