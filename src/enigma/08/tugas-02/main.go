package main

import (
	"bufio"
	f "fmt"
	"os"
	s "strings"
)

type robot struct {
	posX   int
	posY   int
	energi int
}

var perintah string

var bot = robot{
	posX:   0,
	posY:   0,
	energi: 100,
}

func setArah(arah string) {
	arah = s.ToUpper(arah)
	slicePerintah := s.Split(arah, "")
	for _, arah := range slicePerintah {
		switch arah {
		case "T":
			if bot.posX != 10 {
				bot.posX++
				bot.energi -= 5
			} else {
				f.Println("Robot keluar arena")
				break
			}
		case "B":
			if bot.posX != 0 {
				bot.posX--
				bot.energi -= 5
			} else {
				f.Println("Robot keluar arena")
				break
			}
		case "S":
			if bot.posY != 0 {
				bot.posY--
				bot.energi -= 5
			} else {
				f.Println("Robot keluar arena")
				break
			}
		case "U":
			if bot.posY != 10 {
				bot.posY++
				bot.energi -= 5
			} else {
				f.Println("Robot keluar arena")
				break
			}
		}
	}
}

func main() {

	f.Printf("Robot kordinat (%v,%v), energi %v%%\n", bot.posX, bot.posY, bot.energi)
	for {
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
