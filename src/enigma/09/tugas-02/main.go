package main

import (
	"bufio"
	f "fmt"
	"os"
	s "strings"
)

const X_AWAL = 0
const Y_AWAL = 0
const ENERGI_AWAL = 5
const BATAS_X = 10
const BATAS_Y = 10

type robot struct {
	posX   int
	posY   int
	energi int
}

func (b *robot) keTimur() {
	if b.posX != BATAS_X {
		b.posX++
		b.energi -= 5
	} else {
		f.Println("Robot keluar arena")
	}
}

func (b *robot) keBarat() {
	if b.posX != X_AWAL {
		b.posX--
		b.energi -= 5
	} else {
		f.Println("Robot keluar arena")
	}
}

func (b *robot) keSelatan() {
	if b.posY != Y_AWAL {
		b.posY--
		b.energi -= 5
	} else {
		f.Println("Robot keluar arena")
	}
}

func (b *robot) keUtara() {
	if b.posY != BATAS_Y {
		b.posY++
		b.energi -= 5
	} else {
		f.Println("Robot keluar arena")
	}
}

var bot = robot{
	posX:   0,
	posY:   0,
	energi: 100,
}

var perintah string

func setArah(arah string) {
	arah = s.ToUpper(arah)
	slicePerintah := s.Split(arah, "")
	for _, arah := range slicePerintah {
		switch arah {
		case "T":
			bot.keTimur()
		case "B":
			bot.keBarat()
		case "S":
			bot.keSelatan()
		case "U":
			bot.keUtara()
		}
	}
}

func main() {

	f.Printf("Robot kordinat (%v,%v), energi %v%%\n", bot.posX, bot.posY, bot.energi)
	for {
		f.Println("T: timur, B: barat, S: selatan, U: utara")
		f.Printf("Masukkan perintah: ")
		perintah = input()
		if cekPerintah(perintah) {
			setArah(perintah)
			f.Printf("Robot kordinat (%v,%v), energi %v%%\n", bot.posX, bot.posY, bot.energi)
			if bot.energi == 0 {
				f.Println("Energi tidak cukup")
				os.Exit(2)
			}
		} else {
			f.Println("Perintah salah")
		}
	}
}

func cekPerintah(perintah string) bool {
	slicePerintah := s.Split(perintah, "")
	for _, arah := range slicePerintah {
		arah = s.ToUpper(arah)
		if arah == "T" || arah == "B" || arah == "S" || arah == "U" {
			return true
		}
		return false
	}
	return false
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	perintah = scanner.Text()
	return perintah
}
