package main

import "fmt"

type bilBulat int

func (b *bilBulat) kali(pengali int) {
	*b = *b * bilBulat(pengali)
}

type Rectangle struct {
	width  int
	length int
}

type Circle struct {
	radius float32
}

func (rct Rectangle) getLuasKemudianKali(pengali int) int {
	return rct.width * rct.length * pengali
}

func (rct *Rectangle) doubleWidth() {
	rct.width = rct.width * 2
}

func main() {
	rect := Rectangle{
		5,
		10,
	}
	fmt.Println(rect.getLuasKemudianKali(2))
	rect.doubleWidth()
	fmt.Println(rect.getLuasKemudianKali(1))

	rect2 := Rectangle{
		2,
		9,
	}

	fmt.Println(rect2.getLuasKemudianKali(5))

	var b bilBulat
	b = 10
	b.kali(5)
	fmt.Println(b)

	var i int
	i.kali(3)
}
