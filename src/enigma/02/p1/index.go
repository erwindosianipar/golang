package main

import "fmt"

func main() {
	var i int = 97
	var f float32 = 10.5

	var hasil = float32(i) + f
	fmt.Println(hasil)

	var kal string = "ini kalimat"
	fmt.Printf("Hello %s", kal)
	fmt.Println("Hello", kal)
	fmt.Print("Hello", kal)

	var nil int = 100
	fmt.Println(100, "ini juga di print", nil)
	tangkap := fmt.Sprintln("ini akan dicetak")

	fmt.Println("ini semua akan dicetak", tangkap)
}
