package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type hero struct {
	nama   string
	hp     int
	damage int
}

var heros []hero
var nama string
var hp, damage, menu int

func main() {
	if len(heros) != 3 {
		nama, hp, damage = inputHero()
		createHero(nama, hp, damage)
		main()
	} else {
		menu = viewMenu()
		switch menu {
		case 1:
			viewHeroes()
			main()
		case 2:
			attack()
			main()
		case 3:
			healing()
			main()
		case 4:
			exit()
		}
	}
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputan := scanner.Text()
	return inputan
}

func inputHero() (string, int, int) {
	for {
		fmt.Printf("Masukkan nama hero #%v		: ", len(heros)+1)
		nama = input()
		if len(nama) > 2 {
			break
		}
		fmt.Println("Error: nama hero minimal 3 karakter")
	}
	for {
		fmt.Printf("Masukkan HP (darah) hero #%v	: ", len(heros)+1)
		hp, _ = strconv.Atoi(input())
		if hp > 9 && hp < 151 {
			break
		}
		fmt.Println("Error: HP hero minimal 10, maksimal 150")
	}
	for {
		fmt.Printf("Masukkan damage hero #%v		: ", len(heros)+1)
		damage, _ = strconv.Atoi(input())
		if damage > 9 && damage < 10001 {
			break
		}
		fmt.Println("Error: damage hero minimal 10, maksimal 1000")
	}
	fmt.Println()
	return nama, hp, damage
}

func createHero(nama string, hp, damage int) {
	heros = append(heros, hero{
		nama:   nama,
		hp:     hp,
		damage: damage,
	})
}

func viewMenu() int {
	fmt.Printf("%v\n", strings.Repeat("-", 30))
	fmt.Println("1. View all hero")
	fmt.Println("2. Create attack")
	fmt.Println("3. Healing hero")
	fmt.Println("4. Exit the game")
	fmt.Printf("%v\n", strings.Repeat("-", 30))
	for {
		fmt.Printf("Masukkan menu: ")
		menu, _ = strconv.Atoi(input())
		if menu > 0 && menu < 5 {
			break
		}
		fmt.Println("Error: menu yang anda masukkan salah")
	}
	return menu
}

func viewHeroes() {
	for i := 0; i < len(heros); i++ {
		fmt.Println()
		fmt.Println(i + 1)
		fmt.Println("Nama:", heros[i].nama)
		fmt.Println("HP:", heros[i].hp)
		fmt.Println("Damage:", heros[i].damage)
	}
}

func attack() {
	var attacker, toAttack int

	for {
		fmt.Printf("Masukkan hero penyerang: ")
		attacker, _ = strconv.Atoi(input())
		if attacker > 0 && attacker < 4 {
			break
		}
		fmt.Println("Hero hanya (1-3)")
	}

	for {
		fmt.Printf("Masukkan hero untuk diserang: ")
		toAttack, _ = strconv.Atoi(input())
		if toAttack > 0 && toAttack < 4 {
			break
		}
		fmt.Println("Hero hanya (1-3)")
	}

	attacker = attacker - 1
	toAttack = toAttack - 1

	var err = errors.New("")
	if heros[attacker] == heros[toAttack] {
		err = errors.New("Hero tidak bisa menyerang diri sendiri")
		log.Fatal("Fatal error: ", err.Error())
	} else {
		if heros[attacker].hp == 0 {
			err = errors.New("Hero penyerang sudah mati, tidak dapat melakukan serangan")
			log.Fatal("Fatal error: ", err.Error())
		} else {
			if heros[toAttack].hp == 0 {
				err = errors.New("Hero yang ingin diserang sudah mati")
				log.Fatal("Fatal error: ", err.Error())
			} else {
				attackHero(attacker, toAttack)
				viewHeroAttacked(attacker, toAttack)
				main()
			}
		}
	}
}

func healing() {
	var healer, toHeal int

	for {
		fmt.Printf("Masukkan hero healer: ")
		healer, _ = strconv.Atoi(input())
		if healer > 0 && healer < 4 {
			break
		}
		fmt.Println("Hero hanya (1-3)")
	}

	for {
		fmt.Printf("Masukkan hero tujuan: ")
		toHeal, _ = strconv.Atoi(input())
		if toHeal > 0 && toHeal < 4 {
			break
		}
		fmt.Println("Hero hanya (1-3)")
	}

	healer = healer - 1
	toHeal = toHeal - 1

	var err = errors.New("")
	if heros[healer] == heros[toHeal] {
		if heros[healer].hp == 0 {
			err = errors.New("Hero ini sudah mati, tidak bisa melakukan healing")
			log.Fatal("Fatal error: ", err.Error())
		} else {
			healHero(healer, toHeal)
			viewHeroHealed(healer, toHeal)
			main()
		}
	} else {
		if heros[healer].hp == 0 {
			err = errors.New("Hero sudah mati, tidak bisa melakukan healing")
			log.Fatal("Fatal error: ", err.Error())
		} else {
			if heros[toHeal].hp == 0 {
				err = errors.New("Hero tujuan healing sudah mati, tidak dapat melakukan healing")
				log.Fatal("Fatal error: ", err.Error())
			} else {
				healHero(healer, toHeal)
				viewHeroHealed(healer, toHeal)
				main()
			}
		}
	}
}

func exit() {
	fmt.Println("Anda telah keluar dari game")
	os.Exit(0)
}

func healHero(healer, toHeal int) {
	heros[toHeal].hp += heros[healer].damage / 2
}

func viewHeroHealed(healer, toHeal int) {
	fmt.Printf("%v\n", strings.Repeat("-", 30))
	fmt.Printf("HP hero #%v bertambah : %v+\n", toHeal+1, heros[healer].damage/2)
	fmt.Printf("%v\n", strings.Repeat("-", 30))
	fmt.Println()
	fmt.Println(toHeal + 1)
	fmt.Println("Nama:", heros[toHeal].nama)
	fmt.Println("HP:", heros[toHeal].hp)
	fmt.Println("Damage:", heros[toHeal].damage)
}

func attackHero(attacker, toAttack int) {
	if heros[attacker].damage > heros[toAttack].hp {
		heros[toAttack].hp = 0
	} else {
		heros[toAttack].hp -= heros[attacker].damage
	}
}

func viewHeroAttacked(attacker, toAttack int) {
	fmt.Printf("%v\n", strings.Repeat("-", 30))
	fmt.Printf("HP hero #%v berkurang : %v-\n", toAttack+1, heros[attacker].damage)
	fmt.Printf("%v\n", strings.Repeat("-", 30))
	fmt.Println()
	fmt.Println(toAttack + 1)
	fmt.Println("Nama:", heros[toAttack].nama)
	fmt.Println("HP:", heros[toAttack].hp)
	fmt.Println("Damage:", heros[toAttack].damage)
}
