package main

func main() {

	// for i := 1; i <= 10; i++ {
	// 	// fmt.Println("Perulangan ke: ", i)
	// }

	for i := 0; i < 5; i++ {
		b := "*"
		for j := 0; j < i; j++ {
			b += "*"
		}
		println(b)
	}
	/*
	*
	**
	***
	****
	*****

	 */
}
