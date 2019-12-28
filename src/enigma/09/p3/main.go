package main

import "fmt"

type bentuk interface {
	getLuas() float32
	getKeliling() float32
}

type kotak struct {
	panjang float32
	lebar   float32
}

func (k kotak) getLuas() float32 {
	return k.panjang * k.lebar
}

func (k kotak) getKeliling() float32 {
	return (k.panjang + k.lebar) * 2
}

func (k *kotak) tambahLebar() {
	k.lebar++
}

type persegi struct {
	sisi float32
}

func (p persegi) getLuas() float32 {
	return p.sisi * p.sisi
}

func (p persegi) getKeliling() float32 {
	return p.sisi * 4
}

func main() {

	var k = kotak{5, 7}
	fmt.Println("Luas K :", k.getLuas())
	fmt.Println("Keliling K :", k.getKeliling())

	var b bentuk
	b = k
	fmt.Println("Luas B :", b.getLuas())
	fmt.Println("Keliling B :", b.getKeliling())

	var p = persegi{10}
	fmt.Println("Luas P :", p.getLuas())
	fmt.Println("Keliling P :", p.getKeliling())

	b = p
	fmt.Println("Luas B :", b.getLuas())
	fmt.Println("Keliling B :", b.getKeliling())

}
