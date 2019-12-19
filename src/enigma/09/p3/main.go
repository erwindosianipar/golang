package main

import "fmt"

type Bentuk interface {
	getLuas() float32
	getKeliling() float32
}

type Kotak struct {
	panjang float32
	lebar   float32
}

func (k Kotak) getLuas() float32 {
	return k.panjang * k.lebar
}

func (k Kotak) getKeliling() float32 {
	return (k.panjang + k.lebar) * 2
}

func (k *Kotak) tambahLebar() {
	k.lebar++
}

type Persegi struct {
	sisi float32
}

func (p Persegi) getLuas() float32 {
	return p.sisi * p.sisi
}

func (p Persegi) getKeliling() float32 {
	return p.sisi * 4
}

func main() {

	var k = Kotak{5, 7}
	fmt.Println("Luas K :", k.getLuas())
	fmt.Println("Keliling K :", k.getKeliling())

	var b Bentuk
	b = k
	fmt.Println("Luas B :", b.getLuas())
	fmt.Println("Keliling B :", b.getKeliling())

	var p = Persegi{10}
	fmt.Println("Luas P :", p.getLuas())
	fmt.Println("Keliling P :", p.getKeliling())

	b = p
	fmt.Println("Luas B :", b.getLuas())
	fmt.Println("Keliling B :", b.getKeliling())

}
