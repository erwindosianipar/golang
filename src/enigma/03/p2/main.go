package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	j := 0
	for j < 10 {
		println(j)
		j++
	}

	ki := 0
	for {
		if ki < 2 {
			println(ki)
		} else {
			break
		}
		ki++
	}
}
